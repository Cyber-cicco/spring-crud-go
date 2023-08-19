package com.example.springgo.dto;

import com.example.springgo.entites.Thread;
import com.example.springgo.dto.ThreadDto;    

import org.springframework.stereotype.Component;


@Component

public class ThreadTransformer {

    public ThreadDto tothreadDto(Thread entity){
        ThreadDto dto = new ThreadDto();
        dto.setId(entity.getId());
        dto.setMessageList(entity.getMessageList());
        dto.setCreateur(entity.getCreateur());
        dto.setDateCreation(entity.getDateCreation());
        dto.setStatus(entity.getStatus());

        //TODO : implémenter les méthodes pour les champs complexes
        return dto;
    }      

    public Thread tothread(ThreadDto dto){
        Thread entity = new Thread();
        entity.setId(dto.getId());
        entity.setMessageList(dto.getMessageList());
        entity.setCreateur(dto.getCreateur());
        entity.setDateCreation(dto.getDateCreation());
        entity.setStatus(dto.getStatus());

        //TODO : implémenter les méthodes pour les champs complexes
        return entity;
    }      

}
