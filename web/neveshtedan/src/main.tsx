import React from 'react'
import ReactDOM from 'react-dom/client'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import NewWriting from './components/NewWriting.tsx';
import { ApolloClient, ApolloProvider, InMemoryCache } from '@apollo/client';
import './index.css'
import TemplatesPage from './components/Templates.tsx';

const client = new ApolloClient({
  uri: 'http://localhost:8080/query',
  cache: new InMemoryCache(),
});

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <Router>
      <Routes>
        <Route exact path="/" element={<ApolloProvider client={client}><TemplatesPage /></ApolloProvider>} />
        <Route path="/new-page/:id" element={<ApolloProvider client={client}><NewWriting /></ApolloProvider>} />       
      </Routes>
    </Router>
  </React.StrictMode>,
)
