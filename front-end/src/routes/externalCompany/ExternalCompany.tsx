import ExternalCompanyDetails from "../../components/externalCompanies/ExternalCompany";
import { PRIVILEGES } from "../../enums/privileges";
import { hasPrivilege } from "../../helpers/hasPrivileges";
import { useGetLoggedInUserPrivileges } from "../../hooks/useGetLoggedInUserPrivileges";

const ExternalCompanyDetailsPage = () => {
    const userPrivileges = useGetLoggedInUserPrivileges();

    if (!hasPrivilege(userPrivileges, PRIVILEGES.external_company_read)) return null;
    
    return <ExternalCompanyDetails />
}

export default ExternalCompanyDetailsPage;