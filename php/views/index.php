<?php

class IndexView{
  private $model;
  private $controller;

  public function __construct($model, $controller){
    $this->model = $model;
    $this->controller = $controller;
  }

  public function output(){
    $res = array('email' => $this->model->email,
    'token' => $this->model->token);
    return json_encode($res);
  }

}
