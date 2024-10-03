import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import 'bootstrap/dist/css/bootstrap.min.css';
import ApolloProviderComponent from './providers/ApolloProvider.tsx';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
      <ApolloProviderComponent />
  </StrictMode>,
)
