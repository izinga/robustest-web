import Link from "next/link";
import {
  Shield,
  Lock,
  Server,
  Eye,
  FileCheck,
  Users,
  Database,
  Wifi,
  CheckCircle,
  Building2,
} from "lucide-react";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Security & Compliance - RobusTest | Enterprise Security",
  description: "RobusTest is designed for enterprise security requirements. 100% on-premise deployment, air-gapped support, and compliance-ready architecture.",
};

const securityFeatures = [
  {
    icon: Server,
    title: "100% On-Premise Deployment",
    description: "RobusTest runs entirely within your infrastructure. No test data, application binaries, or logs ever leave your network.",
    details: [
      "Software installed on your servers",
      "Devices connected to your network",
      "Database under your control",
      "No external dependencies required",
    ],
  },
  {
    icon: Wifi,
    title: "Air-Gapped Environment Support",
    description: "Deploy RobusTest in completely isolated networks with no internet connectivity. Perfect for classified or sensitive projects.",
    details: [
      "Offline installation packages",
      "Local update distribution",
      "No cloud callbacks or telemetry",
      "Fully functional without internet",
    ],
  },
  {
    icon: Lock,
    title: "Data Sovereignty",
    description: "Your data stays in your country, your data center, your control. Meet data residency requirements without compromise.",
    details: [
      "No data transfer to third parties",
      "Choose your storage location",
      "Full data ownership",
      "GDPR and regional compliance ready",
    ],
  },
  {
    icon: Users,
    title: "Access Control",
    description: "Granular permissions and role-based access control. Integrate with your existing identity management systems.",
    details: [
      "Role-based access control (RBAC)",
      "Project-level permissions",
      "SSO integration ready",
      "Audit logging for all actions",
    ],
  },
  {
    icon: Eye,
    title: "Audit & Monitoring",
    description: "Complete visibility into who did what and when. Export logs to your SIEM for centralized security monitoring.",
    details: [
      "Comprehensive audit trails",
      "User activity logging",
      "Test execution history",
      "API access logs",
    ],
  },
  {
    icon: Database,
    title: "Secure Data Handling",
    description: "Your test data, screenshots, videos, and logs are stored securely on your infrastructure with your encryption standards.",
    details: [
      "Encryption at rest (your keys)",
      "Encrypted data in transit",
      "Secure credential storage",
      "Configurable data retention",
    ],
  },
];

const complianceStandards = [
  {
    name: "SOC 2",
    description: "On-premise deployment gives you full control over SOC 2 compliance requirements.",
  },
  {
    name: "HIPAA",
    description: "Healthcare organizations can maintain HIPAA compliance with data staying on-premise.",
  },
  {
    name: "PCI-DSS",
    description: "Financial services can meet PCI-DSS requirements with isolated device testing environments.",
  },
  {
    name: "GDPR",
    description: "European data residency requirements satisfied with local deployment.",
  },
  {
    name: "ISO 27001",
    description: "Integrates with your existing ISO 27001 certified infrastructure.",
  },
  {
    name: "FedRAMP",
    description: "Government agencies can deploy within FedRAMP authorized boundaries.",
  },
];

const architecturePoints = [
  {
    title: "No Cloud Dependencies",
    description: "RobusTest operates entirely within your network. There are no callbacks to external servers, no cloud services required, and no data leaving your premises.",
  },
  {
    title: "Your Infrastructure",
    description: "Deploy on your own servers — physical or virtual. Use your existing security controls, firewalls, and monitoring systems.",
  },
  {
    title: "Network Isolation",
    description: "Devices can be isolated on separate VLANs. Test applications without exposing your production network to test traffic.",
  },
  {
    title: "Credential Management",
    description: "Integrate with your existing secrets management. No credentials stored in plain text. Support for vault integration.",
  },
];

export default function SecurityPage() {
  return (
    <div className="pt-16">
      {/* Hero Section */}
      <section className="bg-gradient-to-br from-gray-900 via-gray-800 to-gray-900 text-white py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="max-w-3xl">
            <div className="flex items-center gap-3 mb-6">
              <Shield className="text-green-400" size={40} />
              <span className="text-green-400 font-semibold">Enterprise Security</span>
            </div>
            <h1 className="text-4xl md:text-5xl font-bold mb-6">
              Built for Enterprise Security Requirements
            </h1>
            <p className="text-xl text-gray-300 leading-relaxed">
              RobusTest is deployed 100% on your premises. Your data never leaves your network. Your security team maintains full control.
            </p>
          </div>
        </div>
      </section>

      {/* Key Security Message */}
      <section className="py-12 bg-green-50 border-b border-green-100">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex flex-col md:flex-row items-center justify-center gap-8 text-center md:text-left">
            <div className="flex items-center gap-4">
              <CheckCircle className="text-green-600" size={32} />
              <div>
                <p className="font-semibold text-gray-900">No Data Leaves Your Network</p>
                <p className="text-sm text-gray-600">100% on-premise deployment</p>
              </div>
            </div>
            <div className="flex items-center gap-4">
              <CheckCircle className="text-green-600" size={32} />
              <div>
                <p className="font-semibold text-gray-900">Air-Gapped Support</p>
                <p className="text-sm text-gray-600">No internet required</p>
              </div>
            </div>
            <div className="flex items-center gap-4">
              <CheckCircle className="text-green-600" size={32} />
              <div>
                <p className="font-semibold text-gray-900">Your Infrastructure</p>
                <p className="text-sm text-gray-600">Your security controls apply</p>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Security Features */}
      <section className="py-16 md:py-24 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
              Security Features
            </h2>
            <p className="text-xl text-gray-600 max-w-2xl mx-auto">
              Designed from the ground up for enterprise security requirements.
            </p>
          </div>
          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
            {securityFeatures.map((feature) => (
              <div
                key={feature.title}
                className="bg-gray-50 rounded-xl p-6"
              >
                <div className="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center mb-4">
                  <feature.icon className="text-blue-600" size={24} />
                </div>
                <h3 className="text-xl font-semibold text-gray-900 mb-2">
                  {feature.title}
                </h3>
                <p className="text-gray-600 mb-4">{feature.description}</p>
                <ul className="space-y-2">
                  {feature.details.map((detail) => (
                    <li key={detail} className="flex items-center gap-2 text-sm text-gray-600">
                      <CheckCircle className="text-green-500 flex-shrink-0" size={16} />
                      {detail}
                    </li>
                  ))}
                </ul>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Architecture */}
      <section className="py-16 md:py-24 bg-gray-900 text-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold mb-4">
              Secure Architecture
            </h2>
            <p className="text-xl text-gray-400 max-w-2xl mx-auto">
              RobusTest integrates seamlessly with your existing security infrastructure.
            </p>
          </div>
          <div className="grid md:grid-cols-2 gap-8">
            {architecturePoints.map((point) => (
              <div
                key={point.title}
                className="bg-gray-800 rounded-xl p-6"
              >
                <h3 className="text-xl font-semibold mb-3">{point.title}</h3>
                <p className="text-gray-400">{point.description}</p>
              </div>
            ))}
          </div>
        </div>
      </section>
    </div>
  );
}
