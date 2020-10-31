import React from 'react';
import { observer } from 'mobx-react';
import CategoryFilter from '../CategoryFilter/CategoryFilter';
import { createProvider } from '../../store/root';
import '@patternfly/react-core/dist/styles/base.css';

const App: React.FC = observer(() => {
  const Provider = createProvider();

  return (
    <Provider>
      <div className="App">
        <CategoryFilter />
      </div>
    </Provider>
  );
});

export default App;
