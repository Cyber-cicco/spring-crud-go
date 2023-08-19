package com.example.springgo.dto;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

@AllArgsConstructor
@NoArgsConstructor
@Data
@Builder
public class ThreadDto {

    private Long id;
    private List<Message> messageList;
    private Utilisateur createur;
    private LocalDate dateCreation;
    private Status status;


}
