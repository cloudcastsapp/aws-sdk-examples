<?php

require_once(__DIR__.'/../vendor/autoload.php');

use Aws\S3\S3Client;

$s3 = new S3Client([
    'region' => 'us-east-2',
    'version' => 'latest',
]);

$bucket = 'bucket-o-fun';

// $result: ['Location' => 'us-east-2']
$result = $s3->createBucket([
    'Bucket' => $bucket,
]);

$s3->waitUntil('BucketExists', [
    'Bucket'  => $bucket,
    '@waiter' => [
        'delay'       => 3, // Wait 3 seconds between polling
        'maxAttempts' => 10, // Max attempts before failing, total of 30 seconds waited here
    ]
]);