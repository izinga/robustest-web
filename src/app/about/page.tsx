import Link from "next/link";
import { Target, Users, Lightbulb, Heart, MapPin, Award } from "lucide-react";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "About Us - RobusTest | Our Story and Mission",
  description: "Learn about RobusTest, our mission to simplify mobile device lab management for enterprises, and the team behind the platform.",
};

const values = [
  {
    icon: Target,
    title: "Customer Focus",
    description: "We build what enterprises actually need, not what's trendy. Every feature comes from real customer requirements.",
  },
  {
    icon: Users,
    title: "Collaboration",
    description: "We work closely with our customers' teams — QA, DevOps, and Security — to ensure successful deployments.",
  },
  {
    icon: Lightbulb,
    title: "Innovation",
    description: "From Smart TV testing to air-gapped deployments, we solve problems others haven't tackled.",
  },
  {
    icon: Heart,
    title: "Reliability",
    description: "Enterprise software must work. We prioritize stability, security, and long-term support.",
  },
];

const milestones = [
  {
    year: "2016",
    title: "Founded",
    description: "RobusTest started at IIIT Hyderabad with a vision to simplify mobile testing.",
  },
  {
    year: "2018",
    title: "First Enterprise Customer",
    description: "Deployed our first on-premise device lab for a major e-commerce company.",
  },
  {
    year: "2020",
    title: "Smart TV Support",
    description: "Became the first platform to offer comprehensive Roku and Apple TV testing.",
  },
  {
    year: "2022",
    title: "100+ Deployments",
    description: "Reached 100+ enterprise device lab deployments across industries.",
  },
  {
    year: "2024",
    title: "Global Expansion",
    description: "Expanded support to enterprises across North America, Europe, and Asia.",
  },
];

export default function AboutPage() {
  return (
    <div className="pt-16">
      {/* Hero Section */}
      <section className="bg-gradient-to-br from-gray-900 via-gray-800 to-gray-900 text-white py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="max-w-3xl">
            <h1 className="text-4xl md:text-5xl font-bold mb-6">
              About RobusTest
            </h1>
            <p className="text-xl text-gray-300 leading-relaxed">
              We&apos;re on a mission to help enterprises run world-class mobile device labs without the complexity.
            </p>
          </div>
        </div>
      </section>

      {/* Our Story */}
      <section className="py-16 md:py-24 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="max-w-3xl mx-auto">
            <h2 className="text-3xl font-bold text-gray-900 mb-6">Our Story</h2>
            <div className="prose prose-lg text-gray-600 space-y-6">
              <p>
                RobusTest was born out of a simple observation: enterprises were spending enormous amounts of time and money managing device labs, yet still struggling with device reliability, test flakiness, and security concerns.
              </p>
              <p>
                Cloud testing services solved some problems but created others — data leaving the network, unpredictable costs, and limited control. We knew there had to be a better way.
              </p>
              <p>
                Founded at IIIT Hyderabad&apos;s Center for Innovation and Entrepreneurship, RobusTest was built to give enterprises the best of both worlds: the convenience of a managed platform with the security and control of on-premise deployment.
              </p>
              <p>
                Today, RobusTest powers device labs at some of the world&apos;s largest mobile app companies, helping them test across Android, iOS, and Smart TV platforms with confidence.
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* Mission & Vision */}
      <section className="py-16 md:py-24 bg-gray-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="grid md:grid-cols-2 gap-12">
            <div className="bg-white rounded-2xl p-8 shadow-sm">
              <div className="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center mb-6">
                <Target className="text-blue-600" size={24} />
              </div>
              <h3 className="text-2xl font-bold text-gray-900 mb-4">Our Mission</h3>
              <p className="text-gray-600 text-lg leading-relaxed">
                To enable every enterprise to run a world-class mobile device lab without the complexity, cost, and security concerns of traditional approaches.
              </p>
            </div>
            <div className="bg-white rounded-2xl p-8 shadow-sm">
              <div className="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center mb-6">
                <Lightbulb className="text-blue-600" size={24} />
              </div>
              <h3 className="text-2xl font-bold text-gray-900 mb-4">Our Vision</h3>
              <p className="text-gray-600 text-lg leading-relaxed">
                A world where testing on real devices is as simple and reliable as testing in a simulator — but with all the benefits of actual hardware.
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* Our Values */}
      <section className="py-16 md:py-24 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
              Our Values
            </h2>
            <p className="text-xl text-gray-600 max-w-2xl mx-auto">
              The principles that guide everything we do.
            </p>
          </div>
          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-8">
            {values.map((value) => (
              <div key={value.title} className="text-center">
                <div className="w-14 h-14 bg-blue-100 rounded-xl flex items-center justify-center mx-auto mb-4">
                  <value.icon className="text-blue-600" size={28} />
                </div>
                <h3 className="text-xl font-semibold text-gray-900 mb-2">
                  {value.title}
                </h3>
                <p className="text-gray-600">{value.description}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Timeline */}
      <section className="py-16 md:py-24 bg-gray-900 text-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold mb-4">Our Journey</h2>
            <p className="text-xl text-gray-400 max-w-2xl mx-auto">
              Key milestones in the RobusTest story.
            </p>
          </div>
          <div className="max-w-3xl mx-auto">
            <div className="space-y-8">
              {milestones.map((milestone, index) => (
                <div key={milestone.year} className="flex gap-6">
                  <div className="flex flex-col items-center">
                    <div className="w-12 h-12 bg-blue-600 rounded-full flex items-center justify-center font-bold">
                      {milestone.year.slice(-2)}
                    </div>
                    {index < milestones.length - 1 && (
                      <div className="w-0.5 h-full bg-gray-700 mt-2"></div>
                    )}
                  </div>
                  <div className="pb-8">
                    <p className="text-blue-400 font-semibold">{milestone.year}</p>
                    <h3 className="text-xl font-semibold mt-1">{milestone.title}</h3>
                    <p className="text-gray-400 mt-2">{milestone.description}</p>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </section>

      {/* Location */}
      <section className="py-16 md:py-24 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-12">
            <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
              Where We&apos;re Based
            </h2>
          </div>
          <div className="max-w-2xl mx-auto">
            <div className="bg-gray-50 rounded-2xl p-8 text-center">
              <div className="w-14 h-14 bg-blue-100 rounded-xl flex items-center justify-center mx-auto mb-6">
                <MapPin className="text-blue-600" size={28} />
              </div>
              <h3 className="text-xl font-semibold text-gray-900 mb-4">
                Hyderabad, India
              </h3>
              <p className="text-gray-600 mb-4">
                Center for Innovation and Entrepreneurship<br />
                International Institute of Information Technology<br />
                Gachibowli, Hyderabad - 500032
              </p>
              <div className="flex items-center justify-center gap-2 text-blue-600">
                <Award size={20} />
                <span className="font-medium">Incubated at IIIT Hyderabad</span>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Careers Section */}
      <section id="careers" className="py-16 md:py-24 bg-gray-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="max-w-3xl mx-auto text-center">
            <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-6">
              Join Our Team
            </h2>
            <p className="text-xl text-gray-600 mb-8">
              We&apos;re always looking for talented people who are passionate about mobile technology, testing, and building great software.
            </p>
            <div className="bg-white rounded-xl p-8 shadow-sm">
              <p className="text-gray-600 mb-6">
                Interested in working with us? We&apos;d love to hear from you. Send us your resume and tell us why you&apos;d be a great fit.
              </p>
              <a
                href="mailto:careers@robustest.com"
                className="inline-block bg-blue-600 text-white px-8 py-4 rounded-lg font-semibold hover:bg-blue-700 transition-colors"
              >
                careers@robustest.com
              </a>
            </div>
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-16 md:py-24 gradient-primary text-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h2 className="text-3xl md:text-4xl font-bold mb-6">
            Ready to Get Started?
          </h2>
          <p className="text-xl text-blue-100 mb-8 max-w-2xl mx-auto">
            Let&apos;s discuss how RobusTest can help your team build better mobile applications.
          </p>
          <Link
            href="/contact"
            className="inline-block bg-white text-blue-600 px-8 py-4 rounded-lg font-semibold hover:bg-gray-100 transition-colors text-lg"
          >
            Contact Us
          </Link>
        </div>
      </section>
    </div>
  );
}
