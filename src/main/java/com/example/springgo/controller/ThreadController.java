package com.example.springgo.controller;

import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;    
import com.example.springgo.dto.ThreadDto;
import org.springframework.http.ResponseEntity;
import com.example.springgo.service.ThreadService;
import org.springframework.web.bind.annotation.*;

import java.util.List;


@RestController
@RequiredArgsConstructor
@RequestMapping("thread")
public class ThreadController {

   
    private final ThreadService threadService;

    @GetMapping
    public ResponseEntity<List<ThreadDto>> getAllThread(){
        return ResponseEntity.ok(threadService.recuperer());
    }

    @PostMapping
    public ResponseEntity<List<ThreadDto>> saveThread(@RequestBody ThreadDto dto){
        return ResponseEntity.ok(threadService.creer(dto));
    }

    @PutMapping
    public ResponseEntity<List<ThreadDto>> changeThread(@RequestBody ThreadDto dto){
        return ResponseEntity.ok(threadService.changer(dto));
    }

    @DeleteMapping
    public ResponseEntity<List<ThreadDto>> deleteThread(@RequestBody ThreadDto dto){
        return ResponseEntity.ok(threadService.supprimer(dto));
    }


}
