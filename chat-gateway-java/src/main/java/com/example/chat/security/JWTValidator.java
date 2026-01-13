package com.example.chat.security;

import org.springframework.stereotype.Service;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;

@Service
public class JWTValidator {
    private final String secret = "super-secret-dev-key";

    public String validateAndGetUserId(String token) {
        // Claims is ultimately a JSON map and any values can be added to it
        Claims claims = Jwts.parserBuilder()
                            .setSigningKey(secret.getBytes())
                            .build()
                            .parseClaimsJws(token)
                            .getBody();
        return claims.get("userId", String.class);
    }
}
