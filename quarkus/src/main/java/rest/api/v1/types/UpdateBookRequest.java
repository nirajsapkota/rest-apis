package rest.api.v1.types;

import javax.validation.constraints.NotBlank;

public class UpdateBookRequest {
    
    @NotBlank
    public int id;
    
    public String title;
    
    public String author;
    
}
