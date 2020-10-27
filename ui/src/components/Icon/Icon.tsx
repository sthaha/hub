import React from 'react';
import {
  CatIcon,
  CertificateIcon,
  BuildIcon,
  DomainIcon,
  UserIcon,
  IconSize
} from '@patternfly/react-icons';
import './Icon.css';

export enum Icons {
  cat = 'cat',
  certificate = 'certificate',
  user = 'user',
  build = 'build',
  domain = 'domain'
}

export interface IconProps {
  id: Icons | keyof Icons;
  size: IconSize | keyof typeof IconSize;
  label: string;
}

const Icon: React.FC<IconProps> = (props: IconProps) => {
  const { id, size, label } = props;
  switch (id) {
    case Icons.cat:
      return <CatIcon size={size} className="hub-icon" label={label} />;
    case Icons.certificate:
      return <CertificateIcon size={size} className="hub-icon" label={label} />;
    case Icons.user:
      return <UserIcon size={size} className="hub-icon" label={label} />;
    case Icons.build:
      return <BuildIcon size={size} className="hub-icon" label={label} />;
    case Icons.domain:
      return <DomainIcon size={size} className="hub-icon" label={label} />;
  }
  return <div></div>;
};

export default Icon;
