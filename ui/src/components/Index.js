import * as React from 'react';

import CssBaseline from '@mui/material/CssBaseline';
import { ThemeProvider} from '@mui/material/styles';
import AppAppBar from './sub/AppAppBar';
import Hero from './sub/Hero';
import Variables from './sub/variables';


export default function LandingPage() {
  
  let {mode, LPtheme, userLoggedIn, toggleColorMode} = Variables();
  

  return (
    <ThemeProvider theme={LPtheme}>
      <CssBaseline />
      <AppAppBar mode={mode} toggleColorMode={toggleColorMode} user={userLoggedIn} />
      {/* <Hero /> */}
    </ThemeProvider>
  );
}