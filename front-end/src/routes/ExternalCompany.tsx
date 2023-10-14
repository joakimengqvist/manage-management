import { useSelector } from "react-redux";
import ExternalCompanyDetails from "../components/externalCompanies/ExternalCompany";
import { hasPrivilege } from "../helpers/hasPrivileges";
import { State } from "../types/state";

const ExternalCompanyDetailsPage: React.FC = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges);

    if (!hasPrivilege(userPrivileges, 'external_company_read')) return null;
    
    return (
        <div style={{padding: '12px 8px'}}>
            <div style={{padding: '4px'}}>
                <ExternalCompanyDetails />
            </div>
        </div>
    )

}

export default ExternalCompanyDetailsPage;