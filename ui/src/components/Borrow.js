import * as React from 'react';
import CssBaseline from "@mui/material/CssBaseline";
import { ThemeProvider} from "@mui/material/styles";
import AppAppBar from "./sub/AppAppBar";
import Paper from '@mui/material/Paper';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TablePagination from '@mui/material/TablePagination';
import TableRow from '@mui/material/TableRow';
import Variables from './sub/variables';
import { deleteIndRecordUser, viewRecords } from '../apis/handles';
import { Alert, CircularProgress, Snackbar, Container, Button, Typography } from "@mui/material";
import { AxiosError } from 'axios';
import { useParams } from 'react-router-dom';
  

function Borrow() {
  let id = useParams()
    let { mode, LPtheme, userLoggedIn, toggleColorMode } = Variables();
    let role = "user"
    if(userLoggedIn){
      role = window.sessionStorage.getItem("role");
    }

    let columns = [
      { id: 'id', label: 'Record ID', minWidth: 100 },
      { id: 'BooksID', label: 'Book ID', minWidth: 100 },
      {
        id: 'Title',
        label: 'Book Title',
        minWidth: 170
      },
      {
        id: 'Genre',
        label: 'Genre',
        minWidth: 170
      },
      {
        id: 'created_at',
        label: 'Created At',
        minWidth: 170,
        format: (value) => value.slice(0, 10),
      },
      {
        id: 'returned_at',
        label: 'Returned At',
        minWidth: 170,
        format: (value) => value.slice(0,10),
      },
    ];
    if(!id.id && (role==="admin" || role==="librarian")){
      columns.push({
        id: 'UsersID',
        label: 'User ID',
        minWidth: 170,
      })
    }
  const [records, setRecords] = React.useState([]);
  const [loading, setLoading] = React.useState(true);
  const [error, setError] = React.useState("");

  const [open, setOpen] = React.useState(false);
  const [msg, setMsg] = React.useState("");
  const [variant, setVariant] = React.useState("success")
  const handleClick = (msg, variant) => {
    setMsg(msg);
    setOpen(true);
    setVariant(variant);
  };

  const handleClose = (event, reason) => {
    if (reason === 'clickaway') {
      return;
    }
    setOpen(false);
    window.location.reload();
  };

  const returnBook = (recordID) => {
    let status = deleteIndRecordUser(recordID)
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
  }

    const [page, setPage] = React.useState(0);
    const [rowsPerPage, setRowsPerPage] = React.useState(25);
  
    const handleChangePage = (event, newPage) => {
      setPage(newPage);
    };
  
    const handleChangeRowsPerPage = (event) => {
      setRowsPerPage(+event.target.value);
      setPage(0);
    };

    React.useEffect(() => {
      let s = viewRecords()
      if(id.id){
        s = viewRecords(id.id);
      }
          s.then((res) => {
            if (res instanceof Array)
              setRecords(res);
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

      if(error){
        return <Snackbar anchorOrigin = 	{{ vertical: 'top', horizontal: 'right' }}  open={true} autoHideDuration={3000} onClose={() => window.location="/"}>
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
    
      if (loading) {
        return <CircularProgress />;
      }
      let rows = records;
    return (
        <ThemeProvider theme={LPtheme}>
      <CssBaseline />
      <AppAppBar mode={mode} toggleColorMode={toggleColorMode} user={userLoggedIn} />
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
      <Container id="features" sx={{ py: { xs: 8, sm: 16 } }}>
      <Typography variant="h4" align='center' gutterBottom>
            User's Borrowing Record History
        </Typography>
      <Paper sx={{ width: '100%', overflow: 'hidden' }}>
        <TableContainer sx={{ maxHeight: 440 }}>
          <Table stickyHeader aria-label="sticky table">
            <TableHead>
              <TableRow>
                {columns.map((column) => (
                  <TableCell
                    key={column.id}
                    align={column.align}
                    style={{ minWidth: column.minWidth }}
                  >
                    {column.label}
                  </TableCell>
                ))}
              </TableRow>
            </TableHead>
            <TableBody>
              {rows
                .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                .map((row) => {
                  return (
                    <TableRow hover role="checkbox" tabIndex={-1} key={row.code}>
                      {columns.map((column) => {
                        const value = row[column.id];
                        if(value)
                          return (
                            <TableCell key={column.id} align={column.align}>
                              {column.format 
                                ? column.format(value)
                                : value}
                            </TableCell>
                          );
                        else if(column.id === "returned_at")
                        return (
                          <TableCell key={column.id} align={column.align}>
                            <Button variant="outlined" size="small" onClick={() => returnBook(row["id"])}>
                              <span>Return Book</span>
                            </Button>
                          </TableCell>
                        );
                        else return <></>
                      })}
                    </TableRow>
                  );
                })}
            </TableBody>
          </Table>
        </TableContainer>
        <TablePagination
          rowsPerPageOptions={[10, 25, 100]}
          component="div"
          count={rows.length}
          rowsPerPage={rowsPerPage}
          page={page}
          onPageChange={handleChangePage}
          onRowsPerPageChange={handleChangeRowsPerPage}
        />
      </Paper>
      </Container>
    </ThemeProvider>
    );
}

export default Borrow