// eslint-disable-next-line import/no-extraneous-dependencies
import {AsideHeader} from '@gravity-ui/navigation';

const Sidebar = () => {
    return (
        <AsideHeader
            menuItems={[{id: 'reports', title: 'Reports'}]}
            compact={false}
            headerDecoration
        />
    );
};

export default Sidebar;
