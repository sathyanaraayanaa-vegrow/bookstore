import * as React from "react";
import TextField from "@mui/material/TextField";
import Dialog from "@mui/material/Dialog";
import DialogTitle from "@mui/material/DialogTitle";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import DialogActions from "@mui/material/DialogActions";
import Button from "@mui/material/Button";
import Autocomplete, { createFilterOptions } from "@mui/material/Autocomplete";
import { addAuthors, getAuthors } from "../../apis/handles";
import { Alert, CircularProgress, Grid, Snackbar } from "@mui/material";
import { AxiosError } from "axios";

const filter = createFilterOptions();

export default function AuthorsAutoComplete(props) {
const inText = React.useRef(null);
  const [value, setValue] = React.useState(null);
  const [id, setId] = React.useState("");
  const [open, toggleOpen] = React.useState(false);
  const [focus, setfocus] = React.useState(false);

  const [authors, setAuthors] = React.useState([]);
  const [loading, setLoading] = React.useState(true);
  const [error, setError] = React.useState("");

  const [opened, setOpen] = React.useState(false);
  const [msg, setMsg] = React.useState("");
  const [variant, setVariant] = React.useState("success")
  const handleClick = (msg, variant, l) => {
    setMsg(msg);
    setOpen(true);
    setVariant(variant);
  };

  const handledClose = (event, reason) => {
    if (reason === 'clickaway') {
      return;
    }
    setOpen(false);
    setfocus(false);
  };

  const handleClose = () => {
    setfocus(false);
    setDialogValue({
      Name: "",
    });
    toggleOpen(false);
   
  };

  const [dialogValue, setDialogValue] = React.useState({
    Name: "",
  });

  const handleSubmitBt = (event) => {
    event.preventDefault();
    let status = addAuthors(dialogValue.Name);
    setValue({
        Name: dialogValue.Name,
      })
    status.then(res => {
      if (res instanceof AxiosError){
        if(res.hasOwnProperty("response"))
          handleClick(res.message + " | " + res.response.data, "error");
        else
          handleClick(res.message, "error");
      } else {
            setfocus(true)
          setId(res);
      }
      console.log("s", typeof(res), res instanceof AxiosError);
    });
    handleClose();
  };

  React.useEffect(() => {
    getAuthors()
      .then((res) => {
        if (res instanceof Array) {
            setAuthors(res);
            for (let i in res) {
                if(res[i].ID === parseInt(props.id)){
                    setValue(res[i].Name);
                    setId(res[i].ID);
                }
              }
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
  }, [props]);

  if (error) {
    return (
      <Snackbar
        anchorOrigin={{ vertical: "top", horizontal: "right" }}
        open={true}
        autoHideDuration={3000}
        onClose={() => (window.location = "/")}
      >
        <Alert
          onClose={handledClose}
          severity="error"
          variant="filled"
          sx={{ width: "100%" }}
        >
          {error}
        </Alert>
      </Snackbar>
    );
  }

  return (
    <React.Fragment>
        <Snackbar anchorOrigin = 	{{ vertical: 'top', horizontal: 'right' }}  open={opened} autoHideDuration={3000} onClose={handledClose}>
                <Alert
                    onClose={handledClose}
                    severity={variant}
                    variant="filled"
                    sx={{ width: '100%' }}
                >
                    {msg}
                </Alert> 
      </Snackbar>
      <Autocomplete
        value={value}
        onChange={(event, newValue) => {
          if (typeof newValue === "string") {
            // timeout to avoid instant validation of the dialog's form.
            setTimeout(() => {
              toggleOpen(true);
              setDialogValue({
                Name: newValue,
              });
            });
          } else if (newValue && newValue.inputValue) {
            toggleOpen(true);
            setDialogValue({
              Name: newValue.inputValue,
            });
          } else {
            setValue(newValue);
          }
        }}
        filterOptions={(options, params) => {
          const filtered = filter(options, params);

          if (params.inputValue !== "") {
            filtered.push({
              inputValue: params.inputValue,
              Name: `Add "${params.inputValue}"`,
            });
          }

          return filtered;
        }}
        id="authors-autocomplete"
        options={authors}
        getOptionLabel={(option) => {
          // for example value selected with enter, right from the input
          if (typeof option === "string") {
            return option;
          }
          if (option.inputValue) {
            return option.inputValue;
          }
          setId(option.ID);
          return option.Name;
        }}

        selectOnFocus
        clearOnBlur
        handleHomeEndKeys
        renderOption={(props, option) => <li {...props}>{option.Name}</li>}
        sx={{ width: 300 }}
        freeSolo
        renderInput={(params) => (
            <Grid container spacing={2}  mx={2}>
            <Grid item xs={8} >
          <TextField
            {...params}
            label="Author Name"
            name="authorname"
            InputProps={{
              ...params.InputProps,
              endAdornment: (
                <React.Fragment>
                  {loading ? (
                    <CircularProgress color="inherit" size={20} />
                  ) : null}
                  {params.InputProps.endAdornment}
                </React.Fragment>
              ),
            }}
          />
          </Grid>
          <Grid item xs={4}>
          <TextField
            label="Author ID"
            ref={inText}
            name="authorid"
            value={id}
            size="small"
            focused = {focus}
          />
          </Grid>
          </Grid>
        )}
      />
      <Dialog open={open} onClose={handleClose}>
        <form onSubmit={handleSubmitBt}>
          <DialogTitle>Add a new Author</DialogTitle>
          <DialogContent>
            <DialogContentText>
              Did you miss any author in our list? Please, add it!
            </DialogContentText>
            <TextField
              autoFocus
              margin="dense"
              id="name"
              value={dialogValue.Name}
              onChange={(event) =>
                setDialogValue({
                  ...dialogValue,
                  Name: event.target.value,
                })
              }
              label="Name"
              type="text"
              variant="standard"
            />
          </DialogContent>
          <DialogActions>
            <Button onClick={handleClose}>Cancel</Button>
            <Button type="submit">Add</Button>
          </DialogActions>
        </form>
      </Dialog>
    </React.Fragment>
  );
}
