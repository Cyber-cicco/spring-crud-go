package com.example.springgo.repository;

import com.example.springgo.entites.Thread;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Optional;


public interface ThreadRepository extends JpaRepository<Thread, Long>  {

}
