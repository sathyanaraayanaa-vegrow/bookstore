import React from 'react'
import { Alert, Snackbar } from "@mui/material";

function AlertMessage() {
  const [open, setOpen] = React.useState(false);
  const [msg, setMsg] = React.useState();

  const handleClick = (msg) => {
    setMsg(msg);
    setOpen(true);
  };

  const handleClose = (event, reason) => {
    if (reason === 'clickaway') {
      return;
    }
    setOpen(false);
    setMsg("");
  };

  return (
    <Snackbar anchorOrigin = 	{{ vertical: 'top', horizontal: 'right' }}  open={open} autoHideDuration={5000} onClose={handleClose}>
                <Alert
                    onClose={handleClose}
                    severity="success"
                    variant="filled"
                    sx={{ width: '100%' }}
                >
                    {msg}
                </Alert>
      </Snackbar>
  )
}

export default AlertMessage;