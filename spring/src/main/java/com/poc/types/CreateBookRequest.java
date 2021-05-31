package com.poc.types;

import javax.validation.constraints.NotBlank;

public class CreateBookRequest {
    
    @NotBlank
    public String title;
    
    @NotBlank
    public String author;

}