package com.example.springgo.entites;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.Builder;
import lombok.NoArgsConstructor;    

import java.util.List;

@NoArgsConstructor
@AllArgsConstructor
@Data
@Entity
@Builder
public class Thread {

    @Id()
    private Long id;      
    @ManyToMany
    @JoinTable(name="thread_message",
            joinColumns = @JoinColumn(name = "thread_id", referencedColumnName = "id"),
            inverseJoinColumns = @JoinColumn(name = "message_id", referencedColumnName = "id")
    )
    private List<Message> messageList;
    private Utilisateur createur;
    private LocalDate dateCreation;
    private Status status;



}
