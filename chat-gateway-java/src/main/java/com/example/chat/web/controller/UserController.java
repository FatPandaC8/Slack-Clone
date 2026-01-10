package com.example.chat.web.controller;

import java.util.List;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

import com.example.chat.infra.grpc.ChatGrpcClient;
import com.example.chat.web.dto.CreateUserHttpRequest;
import com.example.chat.web.dto.CreateUserResponseView;
import com.example.chat.web.dto.UserView;

@RestController
@RequestMapping("/users")
public class UserController {
    private final ChatGrpcClient grpcClient;

    public UserController(ChatGrpcClient grpcClient) {
        this.grpcClient = grpcClient;
    } 

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public CreateUserResponseView create(@RequestBody CreateUserHttpRequest req) {
        return grpcClient.createUser(
            req.name(),
            req.email(), 
            req.password()
        );
    }

    @GetMapping("/conversations/{conversationId}")
    public List<UserView> list(@PathVariable String conversationId) {
        return grpcClient.listUsers(conversationId);
    }
}
