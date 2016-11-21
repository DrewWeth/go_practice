<?php

$router = new Phroute\RouteCollector(new Phroute\RouteParser);
$router->get('api', apiHandler);
$router->get('/',function(){
  return "Welcome!";
});
$dispatcher = new Phroute\Dispatcher($router);

function processInput($uri){
  // $uri = implode('/', array_slice(explode('/', $_SERVER['REQUEST_URI']), 1));
  $uri = implode('/', explode('/', $_SERVER['REQUEST_URI']));
  return $uri;
}

try {
  $response = $dispatcher->dispatch($_SERVER['REQUEST_METHOD'], processInput($_SERVER['REQUEST_URI']));
} catch (Phroute\Exception\HttpRouteNotFoundException $e) {
  var_dump($e);
  die();
} catch (Phroute\Exception\HttpMethodNotAllowedException $e) {
  var_dump($e);
  die();
}
echo $response;
