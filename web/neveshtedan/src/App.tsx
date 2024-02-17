import { useState,useEffect } from 'react';
import './App.css'
import { ApolloClient, InMemoryCache, gql } from '@apollo/client';
import { Link } from 'react-router-dom';

function App() {
  const [data, setData] = useState<any>()
  const client = new ApolloClient({
    uri: 'http://localhost:8080/query',
    cache: new InMemoryCache(),
  });
  useEffect(()=>{
    client
    .query({
  
      query: gql`
  
        query GetLocations {
  
          enteries {
    fields {
      value
      name
    }
  }
  
        }
  
      `,
  
    })
  
    .then((result) => setData(result.data.enteries)).catch((err)=> console.log(err));
  })

  
console.log(data)
  return (
    <><div>{data?.map((post: any, index: number) => (
      <div key={index}>
        {post && post.fields.map((field: any) => (<div>{field.name}</div>))}
      </div>))}
    </div><div>
        <h1>Main Page</h1>
        <Link to="/new-page/1234">Go to New Page</Link>
      </div></>
  )
}

export default App
