import './App.css';
import Container from '@mui/material/Container'
import routes from './routes/routes';

function App() {
  return (
    <Container>
      {routes()}
    </Container>
  );
}

export default App;
