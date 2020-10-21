import React from 'react';
import { shallow } from 'enzyme';
import Icon from './Icon';

describe('Icon Component', () => {
  it('should render icon for Task Kind Filter', () => {
    const component = shallow(<Icon filterName="Task" iconSize="sm" />);
    expect(component).toMatchSnapshot();
    expect(component.find('BuildIcon[label="Task"]').length).toEqual(1);
  });
  it('should render icon for Pipeline Kind Filter', () => {
    const component = shallow(<Icon filterName="Pipeline" iconSize="md" />);
    expect(component).toMatchSnapshot();
    expect(component.find('DomainIcon[label="Pipeline"]').length).toEqual(1);
  });
  it('should render icon for Official Catalog Filter', () => {
    const component = shallow(<Icon filterName="Official" iconSize="lg" />);
    expect(component).toMatchSnapshot();
    expect(component.find('CatIcon[label="Official"]').length).toEqual(1);
  });
  it('should render icon for Verified Catalog Filter', () => {
    const component = shallow(<Icon filterName="Verified" iconSize="xl" />);
    expect(component).toMatchSnapshot();
    expect(component.find('CertificateIcon[label="Verified"]').length).toEqual(1);
  });
  it('should render icon for Community Catalog Filter', () => {
    const component = shallow(<Icon filterName="Community" iconSize="sm" />);
    expect(component).toMatchSnapshot();
    expect(component.find('UserIcon[label="Community"]').length).toEqual(1);
  });
  it('should render icon for Category Filter', () => {
    const component = shallow(<Icon filterName="CLI" iconSize="sm" />);
    expect(component).toMatchSnapshot();
  });
});
