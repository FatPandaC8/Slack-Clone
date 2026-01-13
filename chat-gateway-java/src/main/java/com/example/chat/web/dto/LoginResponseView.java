package com.example.chat.web.dto;

public record LoginResponseView(
    String UserId, 
    String Name, 
    String Token
) {
    
}
