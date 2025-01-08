import {books, users} from '../apis/base';

export async function signIn(data) {
    try{
        const response = await users().post("/login", {
            "user": {
                "email": data.get('email'),
                "password": data.get('password')
            }
        });
        window.sessionStorage.setItem("role", response.data.status.role);
        window.sessionStorage.setItem("jwt-token", response.data.status.token);
        console.log("logged in", response);
        return response
    } catch (error){
      console.log(error)
      return error
    }
}

export async function signUpUser(data, handle) {
    try{
        const response = await users().post("/signup", {
            "user": {
                "name": data.get("name"),
                "email": data.get("email"),
                "password": data.get("password"),
                "Roles": data.get("role")
            }
        });
        if(handle==="user"){
            window.sessionStorage.setItem("role", response.data.status.role);
            window.sessionStorage.setItem("jwt-token", response.data.status.token);
        }
        console.log("logged in", response);
        return response
    } catch (error){
      console.log(error)
      return error
    }
}

export async function logOut(){
    try{
        const response = users().delete("/logout")
        window.sessionStorage.removeItem("role");
        window.sessionStorage.removeItem("jwt-token");
        console.log(response)
        return response;
    } catch(error) {
        console.log(error)
        return error;
    }
}

export async function viewUsers(){
    try{
        const response = await users().get("/users")
        console.log(response)
        return response.data;
    } catch(error) {
        console.log(error)
        return error;
    }
}

export async function deleteUser(userID){
    try {
        const response = await users().delete("/users/"+userID);
        console.log(response);
        return response
    } catch (error) {
        console.log(error);
        return error;
    }
}

export async function viewBooks(id){

    try {
            let response = await books().get("/books");
        if(id){
            response = await books().get("/books/"+id);
        }
        console.log(response);
        return response.data["List of Books"]
    } catch (error) {
        console.log(error);
        return error;
    }
}

export async function createRecordUser(bookID){
    try {
        const response = await books().post("/records", {
            "bookID": bookID
        });
        console.log(response);
        return response
    } catch (error) {
        console.log(error);
        return error;
    }
}

export async function viewRecords(id){
    try {
        let response = await books().get("/records");
        if(id){
            response = await books().get("/records/"+id);
        }
        console.log(response);
        return response.data["List of Records"];
    } catch (error) {
        console.log(error);
        return error;
    }
}

export async function deleteIndRecordUser(recordID){
    try {
        const response = await books().delete("/record/"+recordID);
        console.log(response);
        return response
    } catch (error) {
        console.log(error);
        return error;
    }
}

export async function deleteRecordUser(bookID){
    try {
        const response = await books().delete("/record",{
            data: {
                "bookID": bookID
            }
        });
        console.log(response);
        return response
    } catch (error) {
        console.log(error);
        return error;
    }
}

export async function createBooks(data){
    try {
        const response = await books().post("/books", {
            "Title" : data.get("title"),
            "Year": parseInt(data.get("year")),
            "Prize": parseFloat(data.get("prize")),
            "Genre": data.get("genre"),
            "Isbn": data.get("isbn"),
            "Author_Id": parseInt(data.get("authorid")),
            "Availability": parseInt(data.get("availability")) 
        });
        console.log(response);
        return response
    } catch (error) {
        console.log(error);
        return error;
    }
}

export async function updateBooks(data, id){
    try {
        const response = await books().put("/books/"+id, {
            "Title" : data.get("title"),
            "Year": parseInt(data.get("year")),
            "Prize": parseFloat(data.get("prize")),
            "Genre": data.get("genre"),
            "Isbn": data.get("isbn"),
            "Author_Id": parseInt(data.get("authorid")),
            "Availability": parseInt(data.get("availability")) 
        });
        console.log(response);
        return response
    } catch (error) {
        console.log(error);
        return error;
    }
}

export async function deleteBooks(id){
    try {
        const response = await books().delete("/books/"+id);
        console.log(response);
        return response
    } catch (error) {
        console.log(error);
        return error;
    }
}

export async function getAuthors(){
    try {
        const response = await books().get("/authors");
        console.log(response);
        return response.data["List of Authors"];
    } catch (error) {
        console.log(error);
        return error;
    }
}

export async function addAuthors(Name){
    try {
        const response = await books().post("/authors", {
            "Name": Name
        });
        console.log(response);
        return response.data.id;
    } catch (error) {
        console.log(error);
        return error;
    }
}