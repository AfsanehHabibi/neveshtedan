import { useQuery, gql } from '@apollo/client';
import { Link } from 'react-router-dom';

const GET_TEMPLATES = gql`
  query GetTemplates {
    templates: templates {
      id
      title
    }
  }
`;

const TemplatesPage = () => {
  const { loading, error, data } = useQuery(GET_TEMPLATES);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  const templates = data.templates;

  return (
    <div>
      <h1>Templates Page</h1>
      <ul>
        {templates.map((template) => (
          <li key={template.id}>
            <Link to={`/new-page/${template.id}`}>{template.title}</Link>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default TemplatesPage;
