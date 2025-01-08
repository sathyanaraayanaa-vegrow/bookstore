import * as React from 'react';
import { createTheme } from '@mui/material/styles';
import getLPTheme from './getLPTheme';

function Variables() {
    const [mode, setMode] = React.useState('light');
  const [userLoggedIn, setLoggedIn] = React.useState(false)


  React.useEffect(() => {
    let m = window.localStorage.getItem("app-theme-mode");
    setMode(() => (m ? m : "light"));
    const token = window.sessionStorage.getItem('jwt-token');
    if (token) {
      setLoggedIn(true);
    } else {
      setLoggedIn(false);
    }
  }, []); 

  const toggleColorMode = () => {
    setMode((prev) => (prev === 'dark' ? 'light' : 'dark'));
    window.localStorage.setItem("app-theme-mode", (mode === 'dark' ? 'light' : 'dark'));
  };

  const LPtheme = createTheme(getLPTheme(mode));

  return {mode, LPtheme, userLoggedIn, toggleColorMode}
    
  
}

export default Variables