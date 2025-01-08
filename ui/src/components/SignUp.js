import * as React from "react";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import CssBaseline from "@mui/material/CssBaseline";
import TextField from "@mui/material/TextField";
import Link from "@mui/material/Link";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import LockOutlinedIcon from "@mui/icons-material/LockOutlined";
import Typography from "@mui/material/Typography";
import Container from "@mui/material/Container";
import { ThemeProvider } from "@mui/material/styles";
import Variables from "./sub/variables";
import { Alert, MenuItem, Snackbar } from "@mui/material";
import { AxiosError } from "axios";
import { signUpUser } from "../apis/handles";
import AppAppBar from "./sub/AppAppBar";

export default function SignUp() {
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
    setloc(l);
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

  const handleSubmit = (event) => {
    event.preventDefault();
    const data = new FormData(event.currentTarget);
    if(data.get("name")==="" || data.get("email")==="" || data.get("password")==="" || data.get("confirmpassword")===""){
      handleClick("Error: Form Fields can't be Blank.", "error", "");
      return
    }
    if(data.get("password")!==data.get("confirmpassword")){
      handleClick("Error: Passwords Doesn't Match. Check your password.", "error", "");
      return
    }
    console.log(data.values().toArray());
    let status = ""
    if (role==="admin")
      status = signUpUser(data, "");
    else
      status = signUpUser(data, "user");
    status.then(res => {
      if (res instanceof AxiosError){
        if(res.hasOwnProperty("response"))
          handleClick(res.message + " | " + res.response.data.status.Msg, "error", "");
        else
          handleClick(res.message, "error", "");
      } else {
          handleClick(res.data.status.Msg, "success", "/");
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
      {role==="admin"?<AppAppBar
        mode={mode}
        toggleColorMode={toggleColorMode}
        user={userLoggedIn}
      />:<></>}
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <Box
          sx={{
            marginTop: 15,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <Avatar sx={{ m: 1, bgcolor: "secondary.main" }}>
            <LockOutlinedIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            {role==="admin"?"Create User":"Sign up"}
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
                  autoComplete="given-name"
                  name="name"
                  required
                  fullWidth
                  id="name"
                  label="Name"
                  autoFocus
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  required
                  fullWidth
                  id="email"
                  label="Email Address"
                  name="email"
                  autoComplete="email"
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  required
                  fullWidth
                  name="password"
                  label="Password"
                  type="password"
                  id="password"
                  autoComplete="new-password"
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  required
                  fullWidth
                  name="confirmpassword"
                  label="Confirm Password"
                  type="password"
                  id="confirmpassword"
                  autoComplete="new-password"
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  id="role"
                  required
                  name="role"
                  select
                  label="Role"
                  defaultValue="user"
                >
                  <MenuItem key="user" value="user">
                    User
                  </MenuItem>

                  {role==="admin" ? <MenuItem key="librarian" value="librarian">
                    Librarian
                  </MenuItem> : <div></div>}
                </TextField>
              </Grid>
            </Grid>
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              {role==="admin"?"Create User":"Sign up"}
            </Button>
            {role==="admin"?<></>:<Grid container justifyContent="flex-end">
              <Grid item>
                <Link href="/signin" variant="body2">
                  Already have an account? Sign in
                </Link>
              </Grid>
            </Grid>}
          </Box>
        </Box>
      </Container>
    </ThemeProvider>
  );
}
