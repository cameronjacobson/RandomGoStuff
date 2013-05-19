<?php

for($a=0;$a<10;$a++){
	$client = stream_socket_client('tcp://127.0.0.1:9201');
	fwrite($client,$argv[1].' '.$x.PHP_EOL);
	echo fgets($client);
	fwrite($client,'close'.PHP_EOL);
	echo fgets($client);
	sleep(1);
}
