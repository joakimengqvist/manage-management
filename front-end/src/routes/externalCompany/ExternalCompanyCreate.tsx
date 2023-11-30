import CreateExternalCompany from '../../components/externalCompanies/CreateExternalCompany';
import { PRIVILEGES } from '../../enums/privileges';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { useGetLoggedInUserPrivileges } from '../../hooks/useGetLoggedInUserPrivileges';

const CreateExternalCompanyPage = () => {
    const userPrivileges = useGetLoggedInUserPrivileges();

    if (!hasPrivilege(userPrivileges, PRIVILEGES.external_company_write)) return null;
    
    return <CreateExternalCompany />
}

export default CreateExternalCompanyPage;