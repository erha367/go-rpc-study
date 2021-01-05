<?php

$host = '127.0.0.1';
$port = '8096';
$conn = fsockopen($host, $port, $errno, $errstr, 3);
if (!$conn) {
    var_dump('链接失败');
    exit;
}

$method = 'Arith.Multiply';
$params = ['A' => 4, 'B' => 250];
$err = fwrite($conn, json_encode(array(
        'method' => $method,
        'params' => [$params],
        'id' => 0,
    )) . "\n");
if ($err === false) {
    var_dump('写入数据失败');
    exit;
}
sleep(3);
stream_set_timeout($conn, 3, 0);
$line = fgets($conn);
var_dump($line);
if ($line === false) {
    var_dump('获取响应数据失败');
    exit;
}
var_dump(json_decode($line, true));
//用完要关闭
fclose($conn);
exit;

# 运行 php jsonrpc_client.php