import { useSelector } from 'react-redux';
import CreateSubProject from '../../components/subProjects/CreateSubProject';
import { PRIVILEGES } from '../../enums/privileges';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { State } from '../../types/state';

const CreateExpense = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges);

    if (!hasPrivilege(userPrivileges, PRIVILEGES.sub_project_write)) return null;
    
    return <CreateSubProject />
}

export default CreateExpense;