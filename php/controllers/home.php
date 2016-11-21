<?php

class HomeController{
  private $model;
  private $db;

  public function __construct($db){
    $this->db = $db;
  }

  public function api(){
    $email = $this->db->adapter->findUserByEmail("drew@gmail.com");
    $model = new User($email);
    $view = new IndexView($model, $this);
    return $view->output();
  }

}
