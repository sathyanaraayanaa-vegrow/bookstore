import * as React from "react";
import PropTypes from "prop-types";

import Box from "@mui/material/Box";
import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Divider from "@mui/material/Divider";
import Typography from "@mui/material/Typography";
import MenuItem from "@mui/material/MenuItem";
import Drawer from "@mui/material/Drawer";
import MenuIcon from "@mui/icons-material/Menu";
import ToggleColorMode from "./ToggleColorMode";
import { logOut} from "../../apis/handles";
import { Alert, Snackbar } from "@mui/material";
import { AxiosError } from "axios";


function returnButton(user, handleSubmit) {
  
  if (!user) {
    return (
      <>
        <Button
          color="primary"
          variant="text"
          size="small"
          component="a"
          href="/signin"
          target="_blank"
          onClick={() => (window.location.href = "/signin")}
        >
          Sign in
        </Button>
        <Button
          color="primary"
          variant="contained"
          size="small"
          component="a"
          href="/signup"
          target="_blank"
          onClick={() => (window.location.href = "/signup")}
        >
          Sign up
        </Button>
      </>
    );
  } else {
    return (
      <>
        <Button
          color="primary"
          variant="contained"
          size="small"
          component="a"
          target="_blank"
          onClick={handleSubmit}
        >
          {" "}
          Log out
        </Button>
      </>
    );
  }
}
function adminClicks(user){
  let role = "user";
  if(user){
    role = window.sessionStorage.getItem("role")
  }
  if(role==="admin"){
  return (
  <>
  <MenuItem onClick={() => window.location = "/users"}>
    <Typography variant="body2" color="text.primary">
      Manage Users
    </Typography>
  </MenuItem>
  </>
  );
}
return <></>
}
function librarianClicks(user){
  let role = "user";
  if(user){
    role = window.sessionStorage.getItem("role")
  }
  if(role==="admin" || role==="librarian"){
  return (
    <MenuItem onClick={() => window.location="/createbook"}>
    <Typography variant="body2" color="text.primary">
      Create Books
    </Typography>
  </MenuItem>
  );
}
return <></>
}
function returnClicks(user) {
  return (
  <>
  {adminClicks(user)}
  {librarianClicks(user)}
  <MenuItem onClick={() => window.location = "/books"}>
  <Typography variant="body2" color="text.primary">
    Books
  </Typography>
</MenuItem>
<MenuItem onClick={() => window.location = "/borrow"}>
  <Typography variant="body2" color="text.primary">
    Borrowing History
  </Typography>
</MenuItem>
</>
  );
}

function AppAppBar({ mode, toggleColorMode, user }) {

  const [open, setOpen] = React.useState(false);

  const toggleDrawer = (newOpen) => () => {
    setOpen(newOpen);
  };

  const [opens, setOpens] = React.useState(false);
  const [msg, setMsg] = React.useState("");
  const [variant, setVariant] = React.useState("success")
  const handleClick = (msg, variant) => {
    setMsg(msg);
    setOpens(true);
    setVariant(variant);
  };

  const handleClose = (event, reason) => {
    if (reason === 'clickaway') {
      return;
    }
    setOpens(false);
    window.location = "/";
  };
  const handleSubmit = (event) => {
    event.preventDefault();
    let status = logOut();
    status.then(res => {
      if (res instanceof AxiosError){
        if(res.hasOwnProperty("response"))
          handleClick(res.message + " | " + res.response.data["Msg"], "error");
        else
          handleClick(res.message, "error");
      } else {
          handleClick(res.data.Msg, "success");
      }
      console.log("s", res, res instanceof AxiosError);
    })
  };

  // const scrollToSection = (sectionId) => {
  //   const sectionElement = document.getElementById(sectionId);
  //   const offset = 128;
  //   if (sectionElement) {
  //     const targetScroll = sectionElement.offsetTop - offset;
  //     sectionElement.scrollIntoView({ behavior: "smooth" });
  //     window.scrollTo({
  //       top: targetScroll,
  //       behavior: "smooth",
  //     });
  //     setOpen(false);
  //   }
  // };

  return (
    <div>
      <Snackbar anchorOrigin = 	{{ vertical: 'top', horizontal: 'right' }}  open={opens} autoHideDuration={3000} onClose={handleClose}>
                <Alert
                    onClose={handleClose}
                    severity={variant}
                    variant="filled"
                    sx={{ width: '100%' }}
                >
                    {msg}
                </Alert> 
      </Snackbar>
      <AppBar
        position="fixed"
        sx={{
          boxShadow: 0,
          bgcolor: "transparent",
          backgroundImage: "none",
          mt: 2,
        }}
      >
        <Container maxWidth="lg">
          <Toolbar
            variant="regular"
            sx={(theme) => ({
              display: "flex",
              alignItems: "center",
              justifyContent: "space-between",
              flexShrink: 0,
              borderRadius: "999px",
              bgcolor:
                theme.palette.mode === "light"
                  ? "rgba(255, 255, 255, 0.4)"
                  : "rgba(0, 0, 0, 0.4)",
              backdropFilter: "blur(24px)",
              maxHeight: 40,
              border: "1px solid",
              borderColor: "divider",
              boxShadow:
                theme.palette.mode === "light"
                  ? `0 0 1px rgba(85, 166, 246, 0.1), 1px 1.5px 2px -1px rgba(85, 166, 246, 0.15), 4px 4px 12px -2.5px rgba(85, 166, 246, 0.15)`
                  : "0 0 1px rgba(2, 31, 59, 0.7), 1px 1.5px 2px -1px rgba(2, 31, 59, 0.65), 4px 4px 12px -2.5px rgba(2, 31, 59, 0.65)",
            })}
          >
            <Box
              sx={{
                flexGrow: 1,
                display: "flex",
                alignItems: "center",
                ml: "-18px",
                px: 0,
              }}
            >
              <Typography component="h1" variant="h4" color="text.primary" onClick={() => window.location="/"}>
                VEGROW BOOKSTORE
              </Typography>
              {/* <img
                src={
                  'https://assets-global.website-files.com/61ed56ae9da9fd7e0ef0a967/61f12e6faf73568658154dae_SitemarkDefault.svg'
                }
                style={logoStyle}
                alt="logo of sitemark"
              /> */}
              <Box sx={{ display: { xs: "none", md: "flex" } }}>
                {returnClicks(user)}
              </Box>
            </Box>
            <Box
              sx={{
                display: { xs: "none", md: "flex" },
                gap: 0.5,
                alignItems: "center",
              }}
            >
              <ToggleColorMode mode={mode} toggleColorMode={toggleColorMode} />

              {returnButton(user, handleSubmit)}
            </Box>
            <Box sx={{ display: { sm: "", md: "none" } }}>
              <Button
                variant="text"
                color="primary"
                aria-label="menu"
                onClick={toggleDrawer(true)}
                sx={{ minWidth: "30px", p: "4px" }}
              >
                <MenuIcon />
              </Button>
              <Drawer anchor="right" open={open} onClose={toggleDrawer(false)}>
                <Box
                  sx={{
                    minWidth: "60dvw",
                    p: 2,
                    backgroundColor: "background.paper",
                    flexGrow: 1,
                  }}
                >
                  <Box
                    sx={{
                      display: "flex",
                      flexDirection: "column",
                      alignItems: "end",
                      flexGrow: 1,
                    }}
                  >
                    <ToggleColorMode
                      mode={mode}
                      toggleColorMode={toggleColorMode}
                    />
                  </Box>
                  {returnClicks(user)}
                  <Divider />
                  <MenuItem>{returnButton(user, handleSubmit)}</MenuItem>
                </Box>
              </Drawer>
            </Box>
          </Toolbar>
        </Container>
      </AppBar>
    </div>
  );
}

AppAppBar.propTypes = {
  mode: PropTypes.oneOf(["dark", "light"]).isRequired,
  toggleColorMode: PropTypes.func.isRequired,
};

export default AppAppBar;
