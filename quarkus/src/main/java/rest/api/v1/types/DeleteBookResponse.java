package rest.api.v1.types;

public class DeleteBookResponse {
    private String message;
    
    public DeleteBookResponse(String message) {
        this.message = message;
    }
    
    public String getMessage() {
        return message;
    }
    
    public void setMessage(String message) {
        this.message = message;
    }
}
