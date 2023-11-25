import { useSelector } from 'react-redux';
import User from '../../components/users/User';
import { PRIVILEGES } from '../../enums/privileges';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { State } from '../../types/state';

const UserDetails: React.FC = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges)

    if (!hasPrivilege(userPrivileges, PRIVILEGES.user_read)) return null;

    return <User />
}

export default UserDetails;