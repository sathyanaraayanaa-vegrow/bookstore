import * as React from "react";

import CssBaseline from "@mui/material/CssBaseline";
import { ThemeProvider } from "@mui/material/styles";
import AppAppBar from "./sub/AppAppBar";
import Variables from "./sub/variables";

import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Card from "@mui/material/Card";
import Chip from "@mui/material/Chip";
import Container from "@mui/material/Container";
import Grid from "@mui/material/Unstable_Grid2";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import MenuBookIcon from "@mui/icons-material/MenuBook";
import { createRecordUser, deleteBooks, deleteRecordUser, viewBooks } from "../apis/handles";
import { Alert, CircularProgress, Snackbar, TextField } from "@mui/material";
import SearchIcon from "@mui/icons-material/Search";
import { AxiosError } from "axios";

function BookView() {
  let { mode, LPtheme, userLoggedIn, toggleColorMode } = Variables();
  let role = "user";
  if (userLoggedIn) {
    role = window.sessionStorage.getItem("role");
  }

  const [books, setBooks] = React.useState([]);
  const [displaybooks, setdisplaybooks] = React.useState([])
  const [loading, setLoading] = React.useState(true);
  const [error, setError] = React.useState("");

  const [open, setOpen] = React.useState(false);
  const [msg, setMsg] = React.useState("");
  const [variant, setVariant] = React.useState("success");
  const handleClick = (msg, variant) => {
    setMsg(msg);
    setOpen(true);
    setVariant(variant);
  };

  const handleClose = (event, reason) => {
    if (reason === "clickaway") {
      return;
    }
    setOpen(false);
    window.location.reload();
  };

  const returnButton = (bookID) => {
    return (
      <Grid container>
          <Grid xs={6} my={1}>
            <Button
              variant="outlined"
              size="small"
              onClick={() => borrowBook(bookID)}
            >
              <span>Borrow Book</span>
            </Button>
          </Grid>
          <Grid xs={6} my={1}>
            <Button
              variant="outlined"
              size="small"
              onClick={() => returnBook(bookID)}
            >
              <span>Return Book</span>
            </Button>
          </Grid>
          {role !== "user" ? (
            <>
            <Grid xs={6} my={1}>
              <Button
                variant="outlined"
                size="small"
                onClick={() => (window.location = "/updatebook/" + bookID)}
              >
                <span>Update Book</span>
              </Button>
            </Grid> <Grid xs={6} my={1}>
              <Button
                variant="outlined"
                size="small"
                onClick={() => deleteBook(bookID)}
              >
                <span>Delete Book</span>
              </Button>
            </Grid>
            </>
          ) : (
            <></>
          )}
      </Grid>
    );
  };

  const borrowBook = (bookID) => {
    console.log(bookID);
    let status = createRecordUser(bookID);
    status.then((res) => {
      if (res instanceof AxiosError) {
        if(res.hasOwnProperty("response"))
          handleClick(res.message + " | " + res.response.data["Msg"], "error");
        else
          handleClick(res.message, "error");
      } else {
        handleClick(res.data.Msg, "success");
      }
      console.log("s", res, res instanceof AxiosError);
    });
  };

  const returnBook = (bookID) => {
    let status = deleteRecordUser(bookID);
    status.then((res) => {
      if (res instanceof AxiosError) {
        if(res.hasOwnProperty("response"))
          handleClick(res.message + " | " + res.response.data["Msg"], "error");
        else
          handleClick(res.message, "error");
      } else {
        handleClick(res.data.Msg, "success");
      }
      console.log("s", res, res instanceof AxiosError);
    });
  };

  const deleteBook = (bookID) => {
    let status = deleteBooks(bookID);
    status.then((res) => {
      if (res instanceof AxiosError) {
        if(res.hasOwnProperty("response"))
          handleClick(res.message + " | " + res.response.data["Msg"], "error");
        else
          handleClick(res.message, "error");
      } else {
        handleClick(res.data.Msg, "success");
      }
      console.log("s", res, res instanceof AxiosError);
    });
  };

  React.useEffect(() => {
    viewBooks()
      .then((res) => {
        if (res instanceof Array){
            let x = res.map((it, i) => {
            let avail =
              it["Availability"] > 0
                ? `Availability: ${it["Availability"]}`
                : "Unavailable";
            return {
              icon: <MenuBookIcon />,
              title: it["title"],
              id: it["id"],
              description: `Book ID: ${it["id"]} | Year: ${it["year"]} | Genre: ${it["genre"]} | ${avail}`,
              imageLight:
                'url("/static/images/templates/templates-images/dash-light.png")',
              imageDark:
                'url("/static/images/templates/templates-images/dash-dark.png")',
            };
          })
          setBooks(x);
          setdisplaybooks(x);
        }
        else {
          if(res.hasOwnProperty("response"))
            setError(res.message + " | " + res.response.data["Msg"]);
          else
            setError(res.message);
        }
        setLoading(false);
      })
      .catch((error) => {
        console.log("Error Fetching", error);
        setLoading(false);
      });
  }, []);

  const [selectedItemIndex, setSelectedItemIndex] = React.useState(0);
  const handleItemClick = (index) => {
    setSelectedItemIndex(index);
  };

  if (error) {
    return (
      <Snackbar
        anchorOrigin={{ vertical: "top", horizontal: "right" }}
        open={true}
        autoHideDuration={3000}
        onClose={() => (window.location = "/")}
      >
        <Alert
          onClose={handleClose}
          severity="error"
          variant="filled"
          sx={{ width: "100%" }}
        >
          {error}
        </Alert>
      </Snackbar>
    );
  }

  if (loading) {
    return <CircularProgress />;
  }

  const items = displaybooks;
  const selectedFeature = items[selectedItemIndex];

  const handleChange = (event) => {
    let l = books.filter((a)=>{
      return a.title.toLowerCase().includes(event.target.value) || a.id.toString().includes(event.target.value)
    })
    setdisplaybooks(l);
  }

  return (
    <ThemeProvider theme={LPtheme}>
      <Snackbar
        anchorOrigin={{ vertical: "top", horizontal: "right" }}
        open={open}
        autoHideDuration={3000}
        onClose={handleClose}
      >
        <Alert
          onClose={handleClose}
          severity={variant}
          variant="filled"
          sx={{ width: "100%" }}
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
      <Container id="features" sx={{ py: { xs: 8, sm: 16 } }}>
        <Grid container spacing={6}>
          <Grid xs={12}>
            <Grid container xs={12}>
              <Grid item xs={4}>
                <Typography component="h2" variant="h4" color="text.primary" my={2}>
                  List of Books
                </Typography>
                </Grid>
              <Grid item m={2}  xs={2}>
                <SearchIcon sx={{ fontSize: 50 }} />
              </Grid>
              <Grid item xs={6} m={2}>
                <TextField id="standard-basic" label="Search" variant="standard" onChange={handleChange} fullWidth/>
              </Grid>
            </Grid>
            {items.length===0? <></> :
            <>
            <Grid
              container
              gap={2}
              sx={{ display: { xs: "auto", sm: "none" } }}
            >
              {items.map(({ title }, index) => (
                <Chip
                  key={index}
                  label={title}
                  onClick={() => handleItemClick(index)}
                  sx={{
                    borderColor: (theme) => {
                      if (theme.palette.mode === "light") {
                        return selectedItemIndex === index
                          ? "primary.light"
                          : "";
                      }
                      return selectedItemIndex === index ? "primary.light" : "";
                    },
                    background: (theme) => {
                      if (theme.palette.mode === "light") {
                        return selectedItemIndex === index ? "none" : "";
                      }
                      return selectedItemIndex === index ? "none" : "";
                    },
                    backgroundColor:
                      selectedItemIndex === index ? "primary.main" : "",
                    "& .MuiChip-label": {
                      color: selectedItemIndex === index ? "#fff" : "",
                    },
                  }}
                />
              ))}
            </Grid>
            <Box
              component={Card}
              variant="outlined"
              sx={{
                display: { xs: "auto", sm: "none" },
                mt: 4,
              }}
            >
              <Box
                sx={{
                  backgroundImage: (theme) =>
                    theme.palette.mode === "light"
                      ? items[selectedItemIndex].imageLight
                      : items[selectedItemIndex].imageDark,
                  backgroundSize: "cover",
                  backgroundPosition: "center",
                  minHeight: 280,
                }}
              />
              <Box sx={{ px: 2, pb: 2 }}>
                <Typography
                  color="text.primary"
                  variant="body2"
                  fontWeight="bold"
                >
                  {selectedFeature.title}
                </Typography>
                <Typography
                  color="text.secondary"
                  variant="body2"
                  sx={{ my: 0.5 }}
                >
                  {selectedFeature.description}
                </Typography>
                {returnButton(selectedFeature.id)}
              </Box>
            </Box>
            <Stack
              direction="row"
              justifyContent="center"
              alignItems="flex-start"
              spacing={2}
              useFlexGap
              flexWrap="wrap"
              md={4}
              sx={{ width: "100%", display: { xs: "none", sm: "flex" } }}
            >
              {items.map(({ icon, title, id, description }, index) => (
                <Card
                  key={index}
                  variant="outlined"
                  onClick={() => handleItemClick(index)}
                  sx={{
                    p: 3,
                    height: "fit-content",
                    width: "49%",
                    background: "none",
                    backgroundColor:
                      selectedItemIndex === index
                        ? "action.selected"
                        : undefined,
                    borderColor: (theme) => {
                      if (theme.palette.mode === "light") {
                        return selectedItemIndex === index
                          ? "primary.light"
                          : "grey.200";
                      }
                      return selectedItemIndex === index
                        ? "primary.dark"
                        : "grey.800";
                    },
                  }}
                >
                  <Box
                    sx={{
                      width: "100%",
                      display: "flex",
                      textAlign: "left",
                      flexDirection: { xs: "column", md: "row" },
                      alignItems: { md: "center" },
                      gap: 2.5,
                    }}
                  >
                    <Box
                      sx={{
                        color: (theme) => {
                          if (theme.palette.mode === "light") {
                            return selectedItemIndex === index
                              ? "primary.main"
                              : "grey.300";
                          }
                          return selectedItemIndex === index
                            ? "primary.main"
                            : "grey.700";
                        },
                      }}
                    >
                      {icon}
                    </Box>
                    <Box sx={{ textTransform: "none" }}>
                      <Typography
                        color="text.primary"
                        variant="body2"
                        fontWeight="bold"
                      >
                        {title}
                      </Typography>
                      <Typography
                        color="text.secondary"
                        variant="body2"
                        sx={{ my: 0.5 }}
                      >
                        {description}
                      </Typography>
                      {returnButton(id)}
                    </Box>
                  </Box>
                </Card>
              ))}
            </Stack>
            </>}
          </Grid>
        </Grid>
      </Container>
    </ThemeProvider>
  );
}

export default BookView;
