<?php


interface DatabaseAdapter{
  public function findUserByEmail($email);
}

class Database{
  public $adapter;

  public function __construct(DatabaseAdapter $adapter){
    $this->adapter = $adapter;
  }
}

class SqlDatabase implements DatabaseAdapter{
  public function __construct(){
  }

  public function findUserByEmail($email){
    if ($email === "drew@gmail.com"){
      return $email;
    }
  }
}

class TestDatabase implements DatabaseAdapter{
  public function __construct(){
  }

  public function findUserByEmail($email){
    return "ANNONYMOUS_USER@gmail.com";
  }
}
