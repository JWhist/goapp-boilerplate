import React, { useState } from 'react';
import './index.css';

function App() {
  const [pingResponse, setPingResponse] = useState(null);
  const [metricsResponse, setMetricsResponse] = useState(null);
  const [loading, setLoading] = useState({});

  const testEndpoint = async (endpoint, setter) => {
    setLoading(prev => ({ ...prev, [endpoint]: true }));
    try {
      const response = await fetch(endpoint);
      const data = endpoint === '/metrics' 
        ? await response.text() 
        : await response.json();
      
      setter({
        status: response.status,
        data: data,
        success: response.ok
      });
    } catch (error) {
      setter({
        status: 'ERROR',
        data: error.message,
        success: false
      });
    } finally {
      setLoading(prev => ({ ...prev, [endpoint]: false }));
    }
  };

  const ResponseDisplay = ({ response, title }) => {
    if (!response) return null;
    
    return (
      <div className="response">
        <div className={response.success ? 'status-success' : 'status-error'}>
          Status: {response.status}
        </div>
        <div style={{ marginTop: '10px' }}>
          <strong>{title}:</strong>
        </div>
        <div style={{ marginTop: '5px' }}>
          {typeof response.data === 'string' 
            ? response.data 
            : JSON.stringify(response.data, null, 2)}
        </div>
      </div>
    );
  };

  return (
    <div>
      <div className="header">
        <h1>🚀 Go API Boilerplate</h1>
        <p>A modern Go backend with React frontend</p>
      </div>
      
      <div className="container">
        <div className="card">
          <h2>Welcome to Your Go API Boilerplate!</h2>
          <p>
            This is a statically built React frontend served by your Go backend. 
            The UI is embedded directly into the Go binary for easy deployment.
          </p>
        </div>

        <div className="grid">
          <div className="card">
            <h3>🔗 API Endpoints</h3>
            <p>Test your API endpoints directly from this interface:</p>
            
            <div className="api-test">
              <h4>Ping Endpoint</h4>
              <p>Test the <code>/ping</code> endpoint</p>
              <button 
                className="button" 
                onClick={() => testEndpoint('/ping', setPingResponse)}
                disabled={loading['/ping']}
              >
                {loading['/ping'] ? 'Testing...' : 'Test /ping'}
              </button>
              <ResponseDisplay response={pingResponse} title="Ping Response" />
            </div>

            <div className="api-test">
              <h4>Metrics Endpoint</h4>
              <p>View Prometheus metrics at <code>/metrics</code></p>
              <button 
                className="button" 
                onClick={() => testEndpoint('/metrics', setMetricsResponse)}
                disabled={loading['/metrics']}
              >
                {loading['/metrics'] ? 'Loading...' : 'View Metrics'}
              </button>
              <ResponseDisplay response={metricsResponse} title="Metrics Response" />
            </div>
          </div>

          <div className="card">
            <h3>📚 Documentation</h3>
            <p>Explore the API documentation and test endpoints:</p>
            <ul>
              <li>
                <a href="/swagger/" target="_blank" rel="noopener noreferrer">
                  🔍 Swagger UI Documentation
                </a>
              </li>
              <li>
                <a href="/metrics" target="_blank" rel="noopener noreferrer">
                  📊 Prometheus Metrics
                </a>
              </li>
            </ul>
          </div>

          <div className="card">
            <h3>🛠️ Development</h3>
            <p>Useful development information:</p>
            <ul>
              <li><strong>Backend:</strong> Go with Cobra CLI</li>
              <li><strong>Frontend:</strong> React (statically built)</li>
              <li><strong>API:</strong> OpenAPI 3.0 specification</li>
              <li><strong>Monitoring:</strong> Prometheus metrics</li>
              <li><strong>Documentation:</strong> Swagger UI</li>
            </ul>
          </div>

          <div className="card">
            <h3>🎯 Next Steps</h3>
            <p>Ideas for extending your boilerplate:</p>
            <ul>
              <li>Add authentication and authorization</li>
              <li>Implement database integration</li>
              <li>Add more API endpoints</li>
              <li>Set up CI/CD pipelines</li>
              <li>Add logging and error handling</li>
              <li>Implement rate limiting</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
