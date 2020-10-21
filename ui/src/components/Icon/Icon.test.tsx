import React from 'react';
import { shallow } from 'enzyme';
import Icon from './Icon';
import { IconName } from '../../icons';

describe('Icon Component', () => {
  it('should render icon for Task Kind Filter', () => {
    const component = shallow(<Icon id={IconName.cat} size="sm" label="Task" />);
    expect(component).toMatchSnapshot();
    expect(component.find('CatIcon[label="Task"]').length).toEqual(1);
  });
  it('should render icon for Pipeline Kind Filter', () => {
    const component = shallow(<Icon id={IconName.build} size="md" label="Pipeline" />);
    expect(component).toMatchSnapshot();
    expect(component.find('BuildIcon[label="Pipeline"]').length).toEqual(1);
  });
  it('should render icon for Official Catalog Filter', () => {
    const component = shallow(<Icon id={IconName.domain} size="lg" label="Official" />);
    expect(component).toMatchSnapshot();
    expect(component.find('DomainIcon[label="Official"]').length).toEqual(1);
  });
  it('should render icon for Verified Catalog Filter', () => {
    const component = shallow(<Icon id={IconName.certificate} size="xl" label="Verified" />);
    expect(component).toMatchSnapshot();
    expect(component.find('CertificateIcon[label="Verified"]').length).toEqual(1);
  });
  it('should render icon for Community Catalog Filter', () => {
    const component = shallow(<Icon id={IconName.user} size="sm" label="Community" />);
    expect(component).toMatchSnapshot();
    expect(component.find('UserIcon[label="Community"]').length).toEqual(1);
  });
  it('should render icon for Category Filter', () => {
    const component = shallow(<Icon id={IconName.none} size="sm" label="CLI" />);
    expect(component).toMatchSnapshot();
  });
});
