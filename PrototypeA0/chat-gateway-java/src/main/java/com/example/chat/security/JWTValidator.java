package com.example.chat.security;

import java.nio.charset.StandardCharsets;

import org.springframework.stereotype.Service;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.security.Keys;

@Service
public class JWTValidator {
    private final String secret = "xsaeslrtjtupzegqbzkuohkotelteuxvqlnmwrhonlrvhyvfterihobznadpjttf";

    public String validateAndGetUserId(String token) {
        // Claims is ultimately a JSON map and any values can be added to it
        Claims claims = Jwts.parserBuilder()
                            .setSigningKey(
                                Keys.hmacShaKeyFor(secret.getBytes(StandardCharsets.UTF_8))
                            )
                            .build()
                            .parseClaimsJws(token)
                            .getBody();
        return claims.get("userId", String.class);
    }
}
