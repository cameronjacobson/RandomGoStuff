<?php

// Usage:  php client.php [randomword]
//
// will run 5 clients sequentially.  To simulate multiple
// clients, just start up multiple client processes

for($a=0;$a<5;$a++){
	$client = stream_socket_client('tcp://127.0.0.1:9201');
	for($x=0;$x<50;$x++){
		fwrite($client,$argv[1].' '.$x.PHP_EOL);
		echo fgets($client);
		usleep(50000);
	}
	fwrite($client,'close'.PHP_EOL);
	echo fgets($client).PHP_EOL;
}
