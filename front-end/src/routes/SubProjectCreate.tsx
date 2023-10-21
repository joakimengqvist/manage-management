import { useSelector } from 'react-redux';
import CreateSubProject from '../components/subProjects/createSubProject';
import { PRIVILEGES } from '../enums/privileges';
import { hasPrivilege } from '../helpers/hasPrivileges';
import { State } from '../types/state';

const CreateExpense: React.FC = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges);

    if (!hasPrivilege(userPrivileges, PRIVILEGES.sub_project_write)) return null;
    
    return (
        <div style={{padding: '12px 8px'}}>
            <div style={{padding: '4px'}}>
                <CreateSubProject />
            </div>
        </div>
    )

}

export default CreateExpense;