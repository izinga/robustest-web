import Link from "next/link";
import {
  Smartphone,
  Monitor,
  Settings,
  BarChart3,
  Tv,
  Cpu,
  Users,
  GitBranch,
  FileText,
  Wifi,
  MapPin,
  Battery,
  Camera,
  Video,
  Terminal,
  Layers,
  Zap,
  RefreshCw,
  CheckCircle,
  Webhook,
  Mail,
  MessageSquare,
} from "lucide-react";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Features - RobusTest | Enterprise Mobile Device Lab Platform",
  description: "Explore RobusTest features: manual testing, test automation, performance monitoring, Smart TV testing, CI/CD integration, and more.",
};

const mainFeatures = [
  {
    icon: Smartphone,
    title: "Manual Testing",
    description: "Stream real device screens to your browser with ultra-low latency. Touch, swipe, type, and debug as if the device was in your hands.",
    image: "https://images.unsplash.com/photo-1512941937669-90a1b58e7e9c?w=600&h=400&fit=crop",
    capabilities: [
      "Real-time device streaming via VNC/WebRTC",
      "Touch and gesture input",
      "Keyboard input support",
      "Screenshot and video recording",
      "Device logs in real-time",
      "Multiple simultaneous sessions",
    ],
  },
  {
    icon: Settings,
    title: "Test Automation",
    description: "Run your existing test suites on real devices. Support for all major frameworks with parallel execution across your device fleet.",
    image: "https://images.unsplash.com/photo-1555949963-aa79dcee981c?w=600&h=400&fit=crop",
    capabilities: [
      "Appium integration",
      "Espresso (Android native)",
      "XCUITest (iOS native)",
      "Selenium Grid compatible",
      "Parallel test execution",
      "Test result aggregation",
    ],
  },
  {
    icon: BarChart3,
    title: "Performance Monitoring",
    description: "Built-in performance metrics collection. No additional tools or SDKs required — just run your tests and get insights.",
    image: "https://images.unsplash.com/photo-1551288049-bebda4e38f71?w=600&h=400&fit=crop",
    capabilities: [
      "CPU usage tracking",
      "Memory consumption",
      "Network data (upload/download)",
      "Battery drain analysis",
      "Frame rate monitoring",
      "Historical trend analysis",
    ],
  },
  {
    icon: Tv,
    title: "Smart TV Testing",
    description: "The only platform with comprehensive Smart TV support. Test streaming apps on the devices your users actually use.",
    image: "https://images.unsplash.com/photo-1593359677879-a4bb92f829d1?w=600&h=400&fit=crop",
    capabilities: [
      "Roku device support",
      "Apple TV support",
      "Chromecast testing",
      "Remote control simulation",
      "Screen capture and streaming",
    ],
  },
  {
    icon: Cpu,
    title: "CI/CD Integration",
    description: "Fits into your existing development workflow. Trigger tests on every build and get results in your pipeline.",
    image: "https://images.unsplash.com/photo-1618401471353-b98afee0b2eb?w=600&h=400&fit=crop",
    capabilities: [
      "Jenkins plugin",
      "GitHub Actions support",
      "GitLab CI integration",
      "REST API for custom integrations",
      "Webhook notifications",
      "Build artifact management",
    ],
  },
  {
    icon: Users,
    title: "Team Collaboration",
    description: "Built for enterprise teams. Manage projects, control access, and share results across your organization.",
    image: "https://images.unsplash.com/photo-1522071820081-009f0129c71c?w=600&h=400&fit=crop",
    capabilities: [
      "Role-based access control",
      "Project-based organization",
      "Shared device pools",
      "Team activity logs",
      "User management",
      "SSO integration ready",
    ],
  },
];

const additionalFeatures = [
  {
    icon: FileText,
    title: "Detailed Reporting",
    description: "Comprehensive test reports with screenshots, videos, and logs. Export to PDF or integrate with your existing tools.",
  },
  {
    icon: GitBranch,
    title: "Version Control Integration",
    description: "Link test results to commits and branches. Track quality metrics across your release cycle.",
  },
  {
    icon: Wifi,
    title: "Network Simulation",
    description: "Test under various network conditions. Simulate 3G, 4G, or custom bandwidth limitations.",
  },
  {
    icon: MapPin,
    title: "GPS Mocking",
    description: "Set custom GPS coordinates for location-based testing without physically moving devices.",
  },
  {
    icon: Battery,
    title: "Battery Simulation",
    description: "Test app behavior at different battery levels and charging states.",
  },
  {
    icon: Camera,
    title: "Screenshot Captures",
    description: "Capture screenshots during test sessions for debugging, documentation, and bug reporting.",
  },
  {
    icon: Video,
    title: "Video Recording",
    description: "Record test sessions for debugging and documentation. Shareable video links for bug reports.",
  },
  {
    icon: Terminal,
    title: "Device Shell Access",
    description: "Execute ADB commands directly. Full shell access for advanced debugging scenarios.",
  },
  {
    icon: Layers,
    title: "App Management",
    description: "Install, uninstall, and manage app versions across your device fleet. Support for APK, IPA, and TV app packages.",
  },
  {
    icon: Zap,
    title: "Parallel Execution",
    description: "Run tests on multiple devices simultaneously. Reduce test suite execution time dramatically.",
  },
  {
    icon: RefreshCw,
    title: "Device Health Monitoring",
    description: "Automated device health checks. Get alerts when devices need attention or go offline.",
  },
  {
    icon: Monitor,
    title: "Multi-Device Control",
    description: "Control multiple devices from a single interface. Compare app behavior across devices side by side.",
  },
];

const integrations = [
  { name: "Webhooks", icon: Webhook },
  { name: "Email", icon: Mail },
  { name: "Slack", icon: MessageSquare },
  { name: "Chat Message", icon: MessageSquare },
];

export default function FeaturesPage() {
  return (
    <div className="pt-16">
      {/* Hero Section */}
      <section className="bg-gradient-to-br from-gray-900 via-gray-800 to-gray-900 text-white py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="max-w-3xl">
            <h1 className="text-4xl md:text-5xl font-bold mb-6">
              Platform Features
            </h1>
            <p className="text-xl text-gray-300 leading-relaxed">
              Everything your team needs to build, test, and ship quality mobile applications. All deployed on your premises.
            </p>
          </div>
        </div>
      </section>

      {/* Main Features */}
      <section className="py-16 md:py-24 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="space-y-24">
            {mainFeatures.map((feature, index) => (
              <div
                key={feature.title}
                className={`grid md:grid-cols-2 gap-12 items-center ${
                  index % 2 === 1 ? "md:flex-row-reverse" : ""
                }`}
              >
                <div className={index % 2 === 1 ? "md:order-2" : ""}>
                  <div className="flex items-center gap-4 mb-6">
                    <div className="w-14 h-14 bg-blue-100 rounded-xl flex items-center justify-center">
                      <feature.icon className="text-blue-600" size={28} />
                    </div>
                    <h2 className="text-3xl font-bold text-gray-900">
                      {feature.title}
                    </h2>
                  </div>
                  <p className="text-lg text-gray-600 mb-6 leading-relaxed">
                    {feature.description}
                  </p>
                  <ul className="space-y-3">
                    {feature.capabilities.map((capability) => (
                      <li key={capability} className="flex items-center gap-3">
                        <CheckCircle className="text-green-500 flex-shrink-0" size={20} />
                        <span className="text-gray-700">{capability}</span>
                      </li>
                    ))}
                  </ul>
                </div>
                <div className={`rounded-2xl h-80 overflow-hidden shadow-lg ${index % 2 === 1 ? "md:order-1" : ""}`}>
                  <img
                    src={feature.image}
                    alt={feature.title}
                    className="w-full h-full object-cover"
                  />
                </div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Additional Features Grid */}
      <section className="py-16 md:py-24 bg-gray-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
              More Capabilities
            </h2>
            <p className="text-xl text-gray-600 max-w-2xl mx-auto">
              A comprehensive toolkit for every testing scenario.
            </p>
          </div>
          <div className="grid md:grid-cols-3 lg:grid-cols-4 gap-6">
            {additionalFeatures.map((feature) => (
              <div
                key={feature.title}
                className="bg-white rounded-xl p-6 card-hover"
              >
                <div className="w-10 h-10 bg-blue-100 rounded-lg flex items-center justify-center mb-4">
                  <feature.icon className="text-blue-600" size={20} />
                </div>
                <h3 className="text-lg font-semibold text-gray-900 mb-2">
                  {feature.title}
                </h3>
                <p className="text-sm text-gray-600">
                  {feature.description}
                </p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Integrations Section */}
      <section id="integrations" className="py-16 md:py-24 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
              Integration Triggers
            </h2>
            <p className="text-xl text-gray-600 max-w-2xl mx-auto">
              Works with the tools your team already uses.
            </p>
          </div>
          <div className="grid grid-cols-2 md:grid-cols-4 gap-4 max-w-2xl mx-auto">
            {integrations.map((integration) => (
              <div
                key={integration.name}
                className="bg-gray-50 border border-gray-200 rounded-lg px-6 py-4 text-center"
              >
                <integration.icon className="text-blue-600 mx-auto mb-2" size={24} />
                <p className="font-medium text-gray-700">{integration.name}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-16 md:py-24 gradient-primary text-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h2 className="text-3xl md:text-4xl font-bold mb-6">
            See RobusTest in Action
          </h2>
          <p className="text-xl text-blue-100 mb-8 max-w-2xl mx-auto">
            Schedule a demo to see how these features can transform your testing workflow.
          </p>
          <Link
            href="/contact"
            className="inline-block bg-white text-blue-600 px-8 py-4 rounded-lg font-semibold hover:bg-gray-100 transition-colors text-lg"
          >
            Request a Demo
          </Link>
        </div>
      </section>
    </div>
  );
}
