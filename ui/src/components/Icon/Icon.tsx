import React from 'react';
import {
  CatIcon,
  CertificateIcon,
  BuildIcon,
  DomainIcon,
  UserIcon,
  IconSize
} from '@patternfly/react-icons';

import { IconName } from '../../icons';
import './Icon.css';

export interface IconProps {
  id: IconName | keyof IconName;
  size: IconSize | keyof typeof IconSize;
  label: string;
}

const Icon: React.FC<IconProps> = (props: IconProps) => {
  const { id, size, label } = props;
  switch (id) {
    case IconName.cat:
      return <CatIcon size={size} className="hub-icon" label={label} />;
    case IconName.certificate:
      return <CertificateIcon size={size} className="hub-icon" label={label} />;
    case IconName.user:
      return <UserIcon size={size} className="hub-icon" label={label} />;
    case IconName.build:
      return <BuildIcon size={size} className="hub-icon" label={label} />;
    case IconName.domain:
      return <DomainIcon size={size} className="hub-icon" label={label} />;
  }
  return <div></div>;
};

export default Icon;
