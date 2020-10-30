$version='0.0.0'
$tmpContent = Get-Content './tcping.go'
for ($i=0; $i -le $tmpContent.length; $i++)
{
    if($tmpContent[$i] -like '*TCP Ping *')　
    {
        $version=[regex]::Matches($tmpContent[$i],'(([0-9]|([1-9]([0-9]*))).){2}([0-9]|([1-9]([0-9]*)))').Value;
    }
}
$os='windows','linux';
$arch='amd64','386','arm';

foreach($o in $os){
    $env:GOOS=$o;
    foreach($a in $arch){
        if(!($o -eq 'windows' -and $a -eq 'arm')){
            $env:GOARCH=$a;
            $name="tcping_"+$version+"_"+$o+"_"+$a;
            "["+$o+"_"+$a+"]"
            go build -ldflags '-w -s' -o $name;
            upx -9 $name;
        }
    }
}