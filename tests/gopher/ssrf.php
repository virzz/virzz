<?php

error_reporting(E_ALL);

$u = $_REQUEST['url'];

file_put_contents("php://stdout",print_r($u,1)."\n");

$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, $u);
#curl_setopt($ch, CURLOPT_FOLLOWLOCATION, 1);
curl_setopt($ch, CURLOPT_HEADER, 0);
#curl_setopt($ch, CURLOPT_PROTOCOLS, CURLPROTO_HTTP | CURLPROTO_HTTPS);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
if( ! $result = curl_exec($ch))
{
    file_put_contents("php://stdout",print_r(curl_error($ch),1)."\n");
}
curl_close($ch);
file_put_contents("php://stdout",print_r($result,1)."\n");

?>