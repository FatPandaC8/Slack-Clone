package com.example.chat.web.controller.user;

import java.util.List;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

import com.example.chat.infra.grpc.ChatGrpcClient;
import com.example.chat.web.dto.CreateUserHttpRequest;
import com.example.chat.web.dto.CreateUserReponseView;
import com.example.chat.web.dto.LoginRequest;
import com.example.chat.web.dto.LoginResponseView;
import com.example.chat.web.dto.UserView;

@RestController
@RequestMapping("/users")
public class UserController {
    private final ChatGrpcClient grpcClient;

    public UserController(ChatGrpcClient grpcClient) {
        this.grpcClient = grpcClient;
    } 

    @PostMapping("/register")
    @ResponseStatus(HttpStatus.CREATED)
        public CreateUserReponseView register(@RequestBody CreateUserHttpRequest body) {

        var user = grpcClient.register(
            body.name(),
            body.email(),
            body.password()
        );

        return new CreateUserReponseView(
            user.id(),
            user.name()
        );
    }

    @PostMapping("/login")
    public LoginResponseView login(@RequestBody LoginRequest body) {
        return grpcClient.login(
            body.email(),
            body.password()
        );
    }

    @GetMapping("/conversations/{conversationId}")
    public List<UserView> list(
        @PathVariable String conversationId,
        @RequestHeader("Authorization") String authHeader
    ) {
        String token = authHeader.replace("Bearer ", "");
        return grpcClient.listUsers(token, conversationId);
    }
}
