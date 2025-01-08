import * as React from "react";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import CssBaseline from "@mui/material/CssBaseline";
import TextField from "@mui/material/TextField";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Container from "@mui/material/Container";
import { ThemeProvider } from "@mui/material/styles";
import Variables from "./sub/variables";
import { Alert, Snackbar } from "@mui/material";
import { AxiosError } from "axios";
import LocalLibraryIcon from '@mui/icons-material/LocalLibrary';
import AppAppBar from "./sub/AppAppBar";
import AuthorsAutoComplete from "./sub/AuthorsAutoComplete";
import { createBooks } from "../apis/handles";

export default function CreateBooks() {
  let { LPtheme, mode, userLoggedIn, toggleColorMode} = Variables();
  let role = "user";
  if(userLoggedIn){
    role = window.sessionStorage.getItem("role");
  }

  const [open, setOpen] = React.useState(false);
  const [msg, setMsg] = React.useState("");
  const [variant, setVariant] = React.useState("success")
  const [loc, setloc] = React.useState("")
  const handleClick = (msg, variant, l) => {
    setMsg(msg);
    setOpen(true);
    setVariant(variant);
    setloc(l)
  };

  const handleClose = (event, reason) => {
    if (reason === 'clickaway') {
      return;
    }
    setOpen(false);
    if(loc){
        window.location = loc;
    }
  };

  if(!userLoggedIn || (role!=="admin" && role!=="librarian")){
    let error = "Request failed with status code 401 | Error: User Not Logged In. LogIn to Continue."
    if(userLoggedIn){
        error = "Request failed with status code 401 | Error: Unauthorized Access."
    }
    return       <Snackbar anchorOrigin = 	{{ vertical: 'top', horizontal: 'right' }}  open={true} autoHideDuration={3000} onClose={()=>window.location = "/"}>
    <Alert
        onClose={handleClose}
        severity="error"
        variant="filled"
        sx={{ width: '100%' }}
    >
        {error}
    </Alert> 
</Snackbar>
  }
  

  const handleSubmit = (event) => {
    event.preventDefault();
    const data = new FormData(event.currentTarget);
    console.log(data.get("hello"))
    let x = data.values().toArray()
    for(var i in x){
        if(x[i]===""){
            handleClick("Error: Form Fields can't be Blank.", "error", "");
            return
        }
    }

    if(isNaN(parseInt(data.get("year")))){
        handleClick("Error: Year should be a Number", "error", "");
            return
    } else if(parseInt(data.get("year")) < 0){
        handleClick("Error: Year cannot be negative", "error", "");
            return
    }
    if(isNaN(parseFloat(data.get("prize")))){
        handleClick("Error: Prize should be a Float", "error", "");
            return
    } else if(parseFloat(data.get("prize")) < 0){
        handleClick("Error: Prize cannot be negative", "error", "");
            return
    }
    if(isNaN(parseInt(data.get("authorid")))){
        handleClick("Error: Author ID should be a Number", "error", "");
            return
    } 
    if(isNaN(parseInt(data.get("availability")))){
        handleClick("Error: Availability should be a Number", "error", "");
            return
    } else if(parseInt(data.get("availability")) < 0){
        handleClick("Error: Availablity cannot be negative", "error", "");
            return
    }

    let status = createBooks(data);
    status.then(res => {
      if (res instanceof AxiosError){
        if(res.hasOwnProperty("response"))
          handleClick(res.message + " | " + res.response.data["Msg"], "error", "");
        else
          handleClick(res.message, "error", "");
      } else {
          handleClick(res.data.Msg, "success", "/");
      }
      console.log("s", res, res instanceof AxiosError);
    })
  };

  return (
    <ThemeProvider theme={LPtheme}>
      <Snackbar anchorOrigin = 	{{ vertical: 'top', horizontal: 'right' }}  open={open} autoHideDuration={3000} onClose={handleClose}>
                <Alert
                    onClose={handleClose}
                    severity={variant}
                    variant="filled"
                    sx={{ width: '100%' }}
                >
                    {msg}
                </Alert> 
      </Snackbar>
      <CssBaseline />
      <AppAppBar
        mode={mode}
        toggleColorMode={toggleColorMode}
        user={userLoggedIn}
      />
      <CssBaseline />
      <Container component="main" maxWidth="xs" >
        <Box
          sx={{
            marginTop: 15,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <Avatar sx={{ m: 1, bgcolor: "secondary.main" }}>
            <LocalLibraryIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            Create Book
          </Typography>
          <Box
            component="form"
            noValidate
            onSubmit={handleSubmit}
            sx={{ mt: 3 }}
          >
            <Grid container spacing={2}>
              <Grid item xs={12}>
                <TextField
                  name="title"
                  required
                  fullWidth
                  id="title"
                  label="Title"
                  autoFocus
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  required
                  fullWidth
                  id="year"
                  label="Publication Year"
                  name="year"
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  required
                  fullWidth
                  id="prize"
                  label="Prize of Book"
                  name="prize"
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  required
                  fullWidth
                  id="genre"
                  label="Genre"
                  name="genre"
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  required
                  fullWidth
                  id="isbn"
                  label="ISBN"
                  name="isbn"
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  required
                  fullWidth
                  id="availability"
                  label="Availability"
                  name="availability"
                />
              </Grid>
              <Grid item xs>
                <AuthorsAutoComplete />
              </Grid>
            </Grid>
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              Create Book
            </Button>
          </Box>
        </Box>
      </Container>
    </ThemeProvider>
  );
}
