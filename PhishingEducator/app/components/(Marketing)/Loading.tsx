import { Spinner } from '@fluentui/react-components';

import LoadingStyles from './Loading.module.scss';

export default function Loading() {
  return (
    <div className={LoadingStyles.Loading}>
      <Spinner
        label="Lädt, bitte warten..."
        labelPosition="below"
      />
    </div>
  )
}