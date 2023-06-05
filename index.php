<?php

$userInput = fgetc(STDIN);
$hash = hash('sha256', $userInput);

echo $hash;