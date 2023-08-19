package com.example.springgo.service;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.validation.annotation.Validated;    
import com.example.springgo.dto.ThreadDto;
import com.example.springgo.dto.ThreadTransformer;
import com.example.springgo.repository.ThreadRepository;

import java.util.List;

@Service
@Validated
@RequiredArgsConstructor
public class ThreadService {

    private final ThreadRepository threadRepository;
    private final ThreadTransformer threadTransformer;

    public List<ThreadDto> supprimer(ThreadDto dto){
        return null;
    };

    public List<ThreadDto> changer(ThreadDto dto){
        return null;
    }

    public List<ThreadDto> recuperer(){
        return null;
    }   

    public List<ThreadDto> creer(ThreadDto dto){
        return null;
    }

}
