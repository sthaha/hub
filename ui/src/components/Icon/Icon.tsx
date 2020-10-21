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

interface FilterName {
  filterName: string;
  iconSize: IconSize | keyof typeof IconSize;
}

const Icon: React.FC<FilterName> = (props: FilterName) => {
  const { filterName, iconSize } = props;
  switch (filterName) {
    // Support Tiers
    case 'Official':
      return <CatIcon size={iconSize} className="Icon" label={filterName} />;
    case 'Verified':
      return <CertificateIcon size={iconSize} className="Icon" label={filterName} />;
    case 'Community':
      return <UserIcon size={iconSize} className="Icon" label={filterName} />;
    // Kinds
    case 'Task':
      return <BuildIcon size={iconSize} className="Icon" label={filterName} />;
    case 'Pipeline':
      return <DomainIcon size={iconSize} className="Icon" label={filterName} />;
    default:
      return <div></div>;
  }
};

export default Icon;
