import CreateSubProject from '../../components/subProjects/CreateSubProject';
import { PRIVILEGES } from '../../enums/privileges';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { useGetLoggedInUserPrivileges } from '../../hooks/useGetLoggedInUserPrivileges';

const CreateExpense = () => {
    const userPrivileges = useGetLoggedInUserPrivileges();

    if (!hasPrivilege(userPrivileges, PRIVILEGES.sub_project_write)) return null;
    
    return <CreateSubProject />
}

export default CreateExpense;