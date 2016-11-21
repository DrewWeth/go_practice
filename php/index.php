<?php

require 'vendor/autoload.php';

require 'application/db.php';
function getDbInstance(){
  return (new Database( new SqlDatabase()));
}

require 'models/user.php';
require 'views/index.php';
require 'controllers/home.php';


require_once 'application/router.php';


function apiHandler(){
  $controller = new HomeController(getDbInstance());
  return $controller->api();
}
